package tests

import (
	"palworld-chan/internal/service/memory"
	"testing"
)

func aTest_Memory_Info(t *testing.T) {
	checker := memory.NewChecker()
	usage, err := checker.GetMemoryUsage()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("usage:%f", usage)
}
