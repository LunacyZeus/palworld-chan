package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"palworld-chan/pkg/utility/utils/monitor"
)

func init() {
	rootCmd.AddCommand(testCMD)
}

var testCMD = &cobra.Command{
	Use:   "test",
	Short: "Run test",
	Long:  "Run test",
	Run: func(cmd *cobra.Command, args []string) {
		RunTest()
	},
}

func RunTest() {
	log.Println("run test")
	processName := "PalServer-Linux"
	// 获取进程信息
	cpuUsage, memoryUsage, upTime, memPercent := monitor.GetProcessInfo(processName)
	log.Printf("[%s] cpuUsage %s,memoryUsage %s,upTime: %s,memPercent:%.2f", processName, cpuUsage, memoryUsage, upTime, memPercent)

	return
}
