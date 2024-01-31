package rcon

import (
	"errors"
	"fmt"
)

func GetServerInfo(endpoint, password string) (serverInfo string, err error) {
	rconClient, err := New(endpoint, password)
	if err != nil {
		err = errors.New(fmt.Sprintf("连接到rcon失败: %v", err))
		return
	}

	serverInfo, err = rconClient.Info()
	if err != nil {
		err = errors.New(fmt.Sprintf("显示在线用户失败: %v", err))
		return
	}

	return
}
