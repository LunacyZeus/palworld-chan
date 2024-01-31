package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/dao"
	"palworld-chan/pkg/utility/rcon"
	"palworld-chan/pkg/utility/utils"
	"strings"
)

func SendBroadCast(c *fiber.Ctx) error { //发送服务器广播
	// 解析 JSON 数据到结构体
	var broadCast models.BroadCastInput
	if err := c.BodyParser(&broadCast); err != nil {
		return err
	}

	message := strings.ReplaceAll(broadCast.BroadCast, " ", "_")

	RconAddress, RconPort, RconPasswd, err := dao.GetRconInfo()
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("连接到rcon失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	result, err := rconClient.Broadcast(message)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("发送广播失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	res := models.Response{
		Code:    200,
		Result:  fiber.Map{"broadCast": result},
		Message: "ok",
		Type:    "success",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func ShowPlayers(c *fiber.Ctx) error { //显示在线用户
	RconAddress, RconPort, RconPasswd, err := dao.GetRconInfo()
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("连接到rcon失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	result, err := rconClient.ShowPlayers()
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("显示在线用户失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	res := models.Response{
		Code:    200,
		Result:  result,
		Message: "ok",
		Type:    "success",
	}
	return c.Status(fiber.StatusOK).JSON(res)

}

func RconInfo(c *fiber.Ctx) error { //获取游戏服务器info信息
	RconAddress, RconPort, RconPasswd, err := dao.GetRconInfo()
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("连接到rcon失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	result, err := rconClient.Info()
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("显示在线用户失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	version, name := utils.ParseServerInfo(result)

	_ = dao.Set(consts.BUCKET, "ServerName", version)
	_ = dao.Set(consts.BUCKET, "ServerVersion", name)

	res := models.Response{
		Code:    200,
		Result:  fiber.Map{"version": version, "name": name},
		Message: "ok",
		Type:    "success",
	}
	return c.Status(fiber.StatusOK).JSON(res)

}

func RestartServer(c *fiber.Ctx) error { //显示在线用户
	seconds := c.Query("seconds", "30")
	message := c.Query("message", "Server_will_reboot_in_30_s")

	RconAddress, RconPort, RconPasswd, err := dao.GetRconInfo()
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s:%s", RconAddress, RconPort)
	password := RconPasswd

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: fmt.Sprintf("连接到rcon失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	result, err := rconClient.Shutdown(seconds, message)
	if err != nil {
		res := models.Response{
			Code:    300,
			Result:  fiber.Map{"result": result},
			Message: fmt.Sprintf("重启服务端失败: %v", err),
			Type:    "error",
		}
		return c.JSON(res)
	}

	res := models.Response{
		Code:    200,
		Result:  result,
		Message: "",
		Type:    "success",
	}
	return c.Status(fiber.StatusOK).JSON(res)

}
