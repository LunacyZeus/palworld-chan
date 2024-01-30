package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
	"palworld-chan/internal/service/dashboard"
	"palworld-chan/pkg/utility/rcon"
	"strings"
)

func GetGameServerInfo(c *fiber.Ctx) error { //服务器状态 获取
	processName := "PalServer-Win64-Test-Cmd.exe"

	// 获取进程信息
	cpuUsage, memoryUsage, upTime := dashboard.GetProcessInfo(processName)

	result := models.GameServerInfoStruct{
		ProcessName:   processName,
		ServerName:    "-",
		ServerVersion: "-",
		MemoryUsage:   memoryUsage,
		CpuUsage:      cpuUsage,
		UpTime:        upTime,
		LastBackUp:    "2024/01/26",
	}

	res := resp.JsonResp(result)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetServerInfo(c *fiber.Ctx) error { //游戏服务端状态 获取
	CpuLoad := dashboard.GetCpuLoad()
	MemPercentage := dashboard.GetMemPercentage()
	hostName, uptime, os := dashboard.GetHostInfo()
	isVirtualized := dashboard.IsVirtualized()

	result := models.ServerInfoStruct{
		HostName:         hostName,
		Os:               os,
		UpTime:           uptime,
		CpuUsage:         CpuLoad,
		MemoryUsage:      MemPercentage,
		Load:             "-",
		DiskUsage:        "-",
		Bandwith:         "-",
		IsVirtualization: isVirtualized,
	}

	res := resp.JsonResp(result)
	return c.Status(fiber.StatusOK).JSON(res)
}

func SendBroadCast(c *fiber.Ctx) error { //发送服务器广播
	// 解析 JSON 数据到结构体
	var broadCast models.BroadCastInput
	if err := c.BodyParser(&broadCast); err != nil {
		return err
	}

	message := strings.ReplaceAll(broadCast.BroadCast, " ", "_")

	endpoint := "127.0.0.1:25575"
	password := "test1234"

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

	err = rconClient.Broadcast(message)
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
		Result:  fiber.Map{"broadCast": "done."},
		Message: "ok",
		Type:    "success",
	}

	return c.JSON(res)
}

func ShowPlayers(c *fiber.Ctx) error { //显示在线用户
	endpoint := "127.0.0.1:25575"
	password := "test1234"

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
	return c.JSON(res)
}
