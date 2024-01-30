package memory

import (
	"fmt"
)

type MemoryChecker struct{}

func NewChecker() MemoryChecker { //内存检测器
	return MemoryChecker{}
}

// GetMemoryUsage 获取内存占用信息
func (c *MemoryChecker) GetMemoryUsage() (memoryUsage float64, err error) {
	memory, err := getMemoryUsage()
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
		return
	}

	fmt.Printf("Total: %v, Free: %v, Used: %v\n", memory.Total, memory.Free, memory.Used)
	memoryUsage = float64(memory.Used) / float64(memory.Total)
	return
}

// Stop 停止
func (c *MemoryChecker) Stop() {

}

// Clean 内存清理
func (cleaner *MemoryChecker) Clean() {
	return
}
