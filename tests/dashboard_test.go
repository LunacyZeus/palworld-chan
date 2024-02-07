package tests

import (
	"palworld-chan/pkg/utility/utils/monitor"
	"testing"
)

func aTest_ServerInfo(t *testing.T) {
	//monitor.GetCpuInfo()
	CpuLoad := monitor.GetCpuLoad()
	t.Logf("CpuLoad %s", CpuLoad)
	MemPercentage := monitor.GetMemPercentage()
	t.Logf("MemPercentage %s", MemPercentage)
	hostName, uptime, os := monitor.GetHostInfo()
	t.Logf("hostName %s,uptime %s,os %s", hostName, uptime, os)
	//monitor.GetDiskInfo()
	//monitor.GetNetInfo()

	processName := "PalServer-Win64-Test-Cmd.exe"
	// 获取进程信息
	cpuUsage, memoryUsage, upTime, _ := monitor.GetProcessInfo(processName)
	t.Logf("[%s] cpuUsage %s,memoryUsage %s,upTime: %s", processName, cpuUsage, memoryUsage, upTime)

}
