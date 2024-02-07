package tests

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func getCPUModel() (string, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	case "linux":
		cmd = exec.Command("cat", "/proc/cpuinfo")
	case "windows":
		cmd = exec.Command("wmic", "cpu", "get", "caption")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "darwin":
		return strings.TrimSpace(string(output)), nil
	case "linux":
		for _, line := range strings.Split(string(output), "\n") {
			if strings.HasPrefix(line, "model name") {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					return strings.TrimSpace(parts[1]), nil
				}
			}
		}
	case "windows":
		lines := strings.Split(string(output), "\n")
		if len(lines) >= 2 {
			return strings.TrimSpace(lines[1]), nil
		}
	}

	return "", fmt.Errorf("failed to extract CPU model information")
}

func Test_CPU_Info(t *testing.T) {
	// Retrieve CPU information
	cpuModel, err := getCPUModel()
	if err != nil {
		fmt.Println("Failed to retrieve CPU model:", err)
		return
	}
	fmt.Println("CPU Model:", cpuModel)
}
