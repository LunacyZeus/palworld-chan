package dao

import (
	"errors"
	"fmt"
	"palworld-chan/pkg/utility/rcon"
)

func RconInfo() (RconAddress, RconPort, RconPasswd string, err error) {
	serverSetting, err := ServerSetting()
	if err != nil {
		return
	}
	if serverSetting.RconAddress == "" || serverSetting.RconPort == "" || serverSetting.RconPasswd == "" {
		err = errors.New("RCON链接信息不能为空")
		return
	}
	RconAddress = serverSetting.RconAddress
	RconPort = serverSetting.RconPort
	RconPasswd = serverSetting.RconPasswd
	return
}

func RconClient() (client *rcon.Client, err error) {
	serverSetting, err := ServerSetting()
	if err != nil {
		return
	}
	if serverSetting.RconAddress == "" || serverSetting.RconPort == "" || serverSetting.RconPasswd == "" {
		err = errors.New("RCON链接信息不能为空")
		return
	}
	RconAddress := serverSetting.RconAddress
	RconPort := serverSetting.RconPort
	RconPasswd := serverSetting.RconPasswd

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	client, err = rcon.New(endpoint, password)
	if err != nil {
		err = errors.New(fmt.Sprintf("连接到rcon失败: %v", err))
	}

	return
}
