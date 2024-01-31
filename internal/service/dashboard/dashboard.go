package dashboard

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"log"
	"strings"
	"time"
)

// cpu info
func GetCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
}

// 获取cpu负载
func GetCpuLoad() string {
	info, err := load.Avg()
	if err != nil {
		return "-"
	}
	//fmt.Printf("cpu load:%v\n", info)
	return fmt.Sprintf("%.2f", info.Load5)
}

// mem info
func GetMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

// mem info
func GetMemPercentage() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return "-"
	}

	// 计算内存使用率
	usagePercentage := float64(memory.Used) / float64(memory.Total) * 100.0

	return fmt.Sprintf("%.2f", usagePercentage)
}

// host info
func GetHostInfo() (hostName string, uptime string, os string) {
	hInfo, _ := host.Info()
	//fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
	hostName = hInfo.Hostname
	uptime = fmt.Sprintf("%d", hInfo.Uptime)
	os = fmt.Sprintf("%s_%s", hInfo.OS, hInfo.PlatformVersion)
	return
}

// 是否虚拟化
func IsVirtualized() string {
	info, err := host.Info()
	if err != nil {
		return "异常"
	}

	// 检查虚拟化信息
	virtualized := info.VirtualizationSystem == ""
	if virtualized {
		return "否"
	}
	return "是"
}

// disk info
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}
func GetNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func GetProcessInfoByName(processName string) (*process.Process, error) {
	// 获取所有运行中的进程
	processList, err := process.Processes()
	if err != nil {
		return nil, err
	}

	// 遍历进程列表，找到目标进程
	for _, proc := range processList {
		name, err := proc.Name()
		if err != nil {
			continue
		}
		//fmt.Println(name)

		// 忽略大小写进行比较
		if strings.EqualFold(name, processName) {
			return proc, nil
		}
	}

	return nil, fmt.Errorf("Process not found: %s", processName)
}

func GetProcessInfo(processName string) (cpuUsage, memoryUsage, upTime string) {
	// 获取进程信息
	proc, err := GetProcessInfoByName(processName)
	if err != nil {
		cpuUsage = "-"
		memoryUsage = "-"
		return
	}

	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		cpuUsage = "-"
	} else {
		cpuUsage = fmt.Sprintf("%.2f%%", cpuPercent)
	}

	memoryInfo, err := proc.MemoryInfo()
	if err != nil {
		memoryUsage = "-"
	} else {
		memoryUsage = fmt.Sprintf("%d GB", memoryInfo.RSS/1024/1024/1024)
	}

	// 获取进程创建时间
	createTime, err := proc.CreateTime()
	if err != nil {
		upTime = "未运行"
	} else {
		log.Println(createTime)
		upTimeNum := time.Since(time.Unix(int64(createTime/1000), 0))
		upTime = fmt.Sprintf("%s", upTimeNum.String())
	}

	return
}
