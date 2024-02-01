package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"palworld-chan/internal/service/dashboard"
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
	cpuUsage, memoryUsage, upTime := dashboard.GetProcessInfo(processName)
	log.Printf("[%s] cpuUsage %s,memoryUsage %s,upTime: %s", processName, cpuUsage, memoryUsage, upTime)

	return
}
