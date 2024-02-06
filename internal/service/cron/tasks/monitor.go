package tasks

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"palworld-chan/internal/service/dao"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/rcon"
	"palworld-chan/pkg/utility/utils/monitor"
	"sync"
	"time"
)

var onceMonitor sync.Once
var lastMonitorTimeInstance time.Time

func LastMonitorTime() time.Time {
	onceMonitor.Do(func() {
		lastMonitorTimeInstance = time.Now()
	})
	return lastMonitorTimeInstance
}

func MonitorServer() (err error) { //监控服务端
	serverSetting, err := dao.ServerSetting()
	if err != nil {
		return
	}
	if serverSetting.ProcessName == "" {
		return
	}

	CheckPeriod, err := convertor.ToInt(serverSetting.CheckPeriod)
	if err != nil {
		return
	}

	now := time.Now()
	interval := now.Sub(LastMonitorTime())
	if int64(interval.Seconds()) >= CheckPeriod {
		// 获取进程信息
		cpuUsage, memoryUsage, upTime, memPercent := monitor.GetProcessInfo(serverSetting.ProcessName)
		logger.Info("每%d秒检测服务端进程占用,CPU占用(%s),内存占用(%s),内存占用率(%.2f),运行时间(%s)", CheckPeriod, cpuUsage, memoryUsage, memPercent, upTime)
		lastMonitorTimeInstance = time.Now() //重置时间

		var MemoryThreshold int64
		MemoryThreshold, err = convertor.ToInt(serverSetting.MemoryThreshold)
		if err != nil {
			return
		}

		if memPercent > float64(MemoryThreshold) {
			var rconClient *rcon.Client
			rconClient, err = dao.RconClient()
			if err != nil {
				return err
			}

			logger.Error("内存占用超过预期:%.2f>%.2f", memPercent, float64(MemoryThreshold))

			message := fmt.Sprintf("Server_will_reboot_in_%s_s1", serverSetting.RestartDelay)

			var result string
			result, err = rconClient.Shutdown(serverSetting.RestartDelay, message)
			if err != nil {
				return
			}
			logger.Error("重启结束->响应(%s)", result)
		}
	}
	return
}
