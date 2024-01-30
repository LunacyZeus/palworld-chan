package util

import (
	"os"
	"os/exec"
	"palworld-chan/pkg/logger"
	"strings"
)

func runUconvLatin(s string) string {
	var out strings.Builder
	cmd := exec.Command("uconv", "-x", "latin")
	cmd.Stdin = strings.NewReader(s)
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		logger.Error("failed to run uconv", "error", err)
		return s
	}

	return out.String()
}

func escapeString(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.TrimSpace(s)

	runes := []rune(s)
	for i := range runes {
		b := []byte(string(runes[i]))

		if len(b) != 1 {
			runes[i] = '*'
		}
	}

	return string(runes)
}
