package dao

import (
	"fmt"
	"palworld-chan/internal/consts"
	"palworld-chan/pkg/utility/rcon"
	"palworld-chan/pkg/utility/utils"
)

func GetServerInfo() (serverVersion, serverName string, err error) {
	RconAddress, RconPort, RconPasswd, err := RconInfo()
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	serverInfo, err := rcon.GetServerInfo(endpoint, password)
	if err != nil {
		return
	}

	serverVersion, serverName = utils.ParseServerInfo(serverInfo)

	_ = Set(consts.BUCKET, "ServerName", serverName)
	_ = Set(consts.BUCKET, "ServerVersion", serverVersion)

	return
}
