package dao

import (
	"errors"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/pkg/utility/utils"
)

func ServerSetting() (updateServerSetting *models.UpdateServerSettingInput, err error) {
	//从持久层解析到结构体
	serverSettingJson, err := Get(consts.BUCKET, "serverSetting")

	if err != nil {
		return
	}

	// 创建 UpdateServerSettingInput 实例
	//var updateServerSetting models.UpdateServerSettingInput

	// 解析 JSON 字符串到结构体
	err = utils.FromJSONString(serverSettingJson, &updateServerSetting)
	if err != nil {
		return
	}

	return
}

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
