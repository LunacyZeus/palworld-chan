package memory

import (
	"github.com/shirou/gopsutil/mem"
)

func getMemoryUsage() (*mem.VirtualMemoryStat, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return memory, nil
}
