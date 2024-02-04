package dao

import (
	"errors"
	"fmt"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/rcon"
	"palworld-chan/pkg/utility/utils"
)

const (
	userKey    string = "Users"
	userPrefix string = "user_"
)

// 添加用户 使用一个有序集合加map 有序集合负责去重 map存数据
func AddUser(player models.OnlinePlayer) (err error) {
	bucket := consts.SORT_BUCKET
	value, err := utils.ToJSONString(player)
	if err != nil {
		return
	}
	score := player.Online

	_ = Set(consts.USER_BUCKET, fmt.Sprintf("%s_%s", userPrefix, player.PlayerUid), value)

	err = ZAdd(bucket, userKey, player.PlayerUid, score)
	if err != nil {
		return
	}
	return
}

// 用户数 统计
func CountUser() (count int, err error) {
	bucket := consts.SORT_BUCKET
	count, err = ZCard(bucket, userKey)
	if err != nil {
		return
	}
	return
}

// 用户数 列表
func ListUser() (players []models.OnlinePlayer, err error) {
	bucket := consts.SORT_BUCKET
	members, err := ZMembers(bucket, userKey)
	if err != nil {
		return
	}

	for _, item := range members {
		player := models.OnlinePlayer{}
		uid := item.value

		value, err := Get(consts.USER_BUCKET, fmt.Sprintf("%s_%s", userPrefix, uid))
		if err != nil {
			logger.Error("从map获取json异常: %v", err)
			continue
		}
		err = utils.FromJSONString(value, &player)
		//log.Println(value)
		if err != nil {
			logger.Error("解析json异常: %v", err)
			continue
		}
		players = append(players, player)
	}
	return
}

func GetUsersFromRcon() (players []models.OnlinePlayer, err error) {
	RconAddress, RconPort, RconPasswd, err := RconInfo()
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		err = errors.New(fmt.Sprintf("连接到rcon失败: %v", err))
		return
	}

	players, err = rconClient.ShowPlayers()
	if err != nil {
		err = errors.New(fmt.Sprintf("显示在线用户失败: %v", err))
		return
	}

	for _, user := range players {
		err = AddUser(user)
		if err != nil {
			logger.Error("添加用户失败: %v", err)
			continue
		}
	}

	return
}
