package tests

import (
	"palworld-chan/internal/service/dashboard"
	"testing"
)

func Test_ServerInfo(t *testing.T) {
	//dashboard.GetCpuInfo()
	CpuLoad := dashboard.GetCpuLoad()
	t.Logf("CpuLoad %s", CpuLoad)
	MemPercentage := dashboard.GetMemPercentage()
	t.Logf("MemPercentage %s", MemPercentage)
	hostName, uptime, os := dashboard.GetHostInfo()
	t.Logf("hostName %s,uptime %s,os %s", hostName, uptime, os)
	//dashboard.GetDiskInfo()
	//dashboard.GetNetInfo()

	processName := "PalServer-Win64-Test-Cmd.exe"
	// 获取进程信息
	cpuUsage, memoryUsage, upTime := dashboard.GetProcessInfo(processName)
	t.Logf("[%s] cpuUsage %s,memoryUsage %s,upTime: %s", processName, cpuUsage, memoryUsage, upTime)

}
