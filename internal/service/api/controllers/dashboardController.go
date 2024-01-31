package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
	"palworld-chan/internal/service/cron/job"
	"palworld-chan/internal/service/dao"
	"palworld-chan/internal/service/dashboard"
	"palworld-chan/pkg/utility/utils"
)

func GetGameServerInfo(c *fiber.Ctx) error { //服务器状态 获取
	serverSetting, err := dao.ServerSetting()
	if err != nil {
		return err
	}
	processName := serverSetting.ProcessName
	if serverSetting.ProcessName == "" {
		processName = "PalServer-Win64-Test-Cmd.exe"
	}

	ServerName, err := dao.Get(consts.BUCKET, "ServerName")
	if err != nil { //ServerName 为空自动请求一次
		ServerName = ""
	}

	ServerVersion, err := dao.Get(consts.BUCKET, "ServerVersion")
	if err != nil {
		ServerVersion = ""
	}

	if ServerName == "" {
		ServerVersion, ServerName, err = dao.GetServerInfo()
		if err != nil {
			ServerName = ""
			ServerVersion = ""
		}
	}

	// 获取进程信息
	cpuUsage, memoryUsage, upTime := dashboard.GetProcessInfo(processName)
	LastBackUp := job.LastBackUpTime().Format("2006-01-02 15:04:05")

	result := models.GameServerInfoStruct{
		ProcessName:   processName,
		ServerName:    ServerName,
		ServerVersion: ServerVersion,
		MemoryUsage:   memoryUsage,
		CpuUsage:      cpuUsage,
		UpTime:        upTime,
		LastBackUp:    LastBackUp,
	}

	res := resp.JsonResp(result)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetServerInfo(c *fiber.Ctx) error {
	//游戏服务端状态 获取
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

func GetBackUpList(c *fiber.Ctx) error { //备份文件列表 获取
	serverSetting, err := dao.ServerSetting()
	if err != nil {
		return err
	}
	if serverSetting.DestDir == "" {
		err = errors.New("未设置备份目录")
		return err
	}
	// 获取目录下的全部文件名及创建日期
	backUpFiles, err := utils.GetBackUpFiles(serverSetting.DestDir)
	if err != nil {
		return err
	}

	res := resp.JsonResp(backUpFiles)
	return c.Status(fiber.StatusOK).JSON(res)
}
