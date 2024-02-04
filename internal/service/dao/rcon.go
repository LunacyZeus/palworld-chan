package dao

import "errors"

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
