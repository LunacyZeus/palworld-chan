package job

import (
	"github.com/duke-git/lancet/v2/convertor"
	"palworld-chan/internal/service/dao"
	"palworld-chan/internal/service/dashboard"
	"palworld-chan/pkg/logger"
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
		cpuUsage, memoryUsage, upTime, memPercent := dashboard.GetProcessInfo(serverSetting.ProcessName)
		logger.Debug("每%d秒检测服务端进程占用,CPU占用(%s),内存占用(%s),内存占用率(%.2f),运行时间(%s)", CheckPeriod, cpuUsage, memoryUsage, memPercent, upTime)
		lastMonitorTimeInstance = time.Now() //重置时间
	}
	return
}
