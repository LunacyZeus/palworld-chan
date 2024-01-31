package utils

import (
	"encoding/json"
	"errors"
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

func ParseServerInfo(input string) (version, name string) {
	// 查找版本信息的起始和结束位置
	versionStart := strings.Index(input, "[")
	versionEnd := strings.Index(input, "]")
	if versionStart != -1 && versionEnd != -1 {
		version = input[versionStart+1 : versionEnd]
	}

	// 查找名称信息的起始位置
	nameStart := strings.Index(input, "] ")
	if nameStart != -1 {
		name = input[nameStart+2:]
	}

	return version, name
}

// toJSONString 将结构体转换为 JSON 字符串
func ToJSONString(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// fromJSONString 将 JSON 字符串解析为结构体
func FromJSONString(jsonString string, v interface{}) error {
	err := json.Unmarshal([]byte(jsonString), v)
	if err != nil {
		return err
	}
	return nil
}

func CheckExist(dir string) bool {
	_, err := os.Stat(dir)

	// 检测目标目录是否存在
	if os.IsNotExist(err) {
		err = errors.New("Error: Destination directory does not exist.")
		return false
	}
	return true
}
