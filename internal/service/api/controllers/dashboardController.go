package controllers

import (
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
	"palworld-chan/internal/service/dao"
	"palworld-chan/internal/service/dashboard"
)

func GetGameServerInfo(c *fiber.Ctx) error { //服务器状态 获取
	processName := "PalServer-Win64-Test-Cmd.exe"

	ServerName, err := dao.Get(consts.BUCKET, "ServerName")
	if err != nil {
		ServerName = ""
	}

	ServerVersion, err := dao.Get(consts.BUCKET, "ServerVersion")
	if err != nil {
		ServerVersion = ""
	}

	// 获取进程信息
	cpuUsage, memoryUsage, upTime := dashboard.GetProcessInfo(processName)

	result := models.GameServerInfoStruct{
		ProcessName:   processName,
		ServerName:    ServerName,
		ServerVersion: ServerVersion,
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
