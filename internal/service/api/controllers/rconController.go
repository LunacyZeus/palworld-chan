package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/dao"
	"strings"
)

func SendBroadCast(c *fiber.Ctx) error { //发送服务器广播
	// 解析 JSON 数据到结构体
	var broadCast models.BroadCastInput
	if err := c.BodyParser(&broadCast); err != nil {
		return err
	}

	message := strings.ReplaceAll(broadCast.BroadCast, " ", "_")

	rconClient, err := dao.RconClient()
	if err != nil {
		return err
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

func ShowPlayers(c *fiber.Ctx) error {
	refresh := c.Query("refresh", "")

	if refresh == "1" {
		//显示在线用户 通过缓存
		_, err := dao.GetUsersFromRcon()
		if err != nil {
			return err
		}
	}

	//显示在线用户 通过缓存
	result, err := dao.ListUser()

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
	serverVersion, serverName, err := dao.GetServerInfo()
	if err != nil {
		return err
	}

	res := models.Response{
		Code:    200,
		Result:  fiber.Map{"version": serverVersion, "name": serverName},
		Message: "ok",
		Type:    "success",
	}
	return c.Status(fiber.StatusOK).JSON(res)

}

func RestartServer(c *fiber.Ctx) error { //显示在线用户
	seconds := c.Query("seconds", "30")
	message := c.Query("message", "Server_will_reboot_in_30_s")

	rconClient, err := dao.RconClient()
	if err != nil {
		return err
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

func KickPlayer(c *fiber.Ctx) error { //踢出用户
	steamId := c.Query("steamId", "")

	rconClient, err := dao.RconClient()
	if err != nil {
		return err
	}

	result, err := rconClient.KickPlayer(steamId)
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

func BanPlayer(c *fiber.Ctx) error { //禁用用户
	steamId := c.Query("steamId", "")

	rconClient, err := dao.RconClient()
	if err != nil {
		return err
	}

	result, err := rconClient.BanPlayer(steamId)
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
