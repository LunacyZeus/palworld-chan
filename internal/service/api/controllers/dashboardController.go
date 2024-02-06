package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
	"palworld-chan/internal/service/cron/tasks"
	"palworld-chan/internal/service/dao"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/utils"
	"palworld-chan/pkg/utility/utils/monitor"
	"path/filepath"
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
	cpuUsage, memoryUsage, upTime, _ := monitor.GetProcessInfo(processName)
	LastBackUp := tasks.LastBackUpTime().Format("2006-01-02 15:04:05")
	onlineCount := tasks.OnlineCount() //获取在线人数

	result := models.GameServerInfoStruct{
		ProcessName:   processName,
		ServerName:    ServerName,
		ServerVersion: ServerVersion,
		MemoryUsage:   memoryUsage,
		CpuUsage:      cpuUsage,
		UpTime:        upTime,
		OnlineCount:   onlineCount,
		LastBackUp:    LastBackUp,
	}

	res := resp.JsonResp(result)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetServerInfo(c *fiber.Ctx) error {
	//游戏服务端状态 获取
	CpuLoad := monitor.GetCpuLoad()
	MemPercentage := monitor.GetMemPercentage()
	hostName, uptime, os := monitor.GetHostInfo()
	isVirtualized := monitor.IsVirtualized()

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

func DownLoadBackUpFile(c *fiber.Ctx) error { //备份文件 下载
	fileName := c.Query("fileName", "")
	if fileName == "" {
		err := errors.New("备份文件不能为空")
		return err
	}
	// 检查后缀名是否为.zip
	if filepath.Ext(fileName) != ".zip" {
		err := errors.New("web安全考虑,只能下载zip格式文件")
		return err
	}

	serverSetting, err := dao.ServerSetting()
	if err != nil {
		return err
	}

	// 组合路径
	filePath := filepath.Join(serverSetting.DestDir, fileName)
	if !utils.FileOrDirExists(filePath) {
		err = errors.New("文件不存在")
		return err
	}
	logger.Debug("下载备份文件: %s", filePath)
	return c.Status(fiber.StatusOK).Download("D:/test/pal_20240201224637.zip")
}

func GetSaveFileList(c *fiber.Ctx) error { //存档文件列表 获取
	serverSetting, err := dao.ServerSetting()
	if err != nil {
		return err
	}

	if serverSetting.SourceDir == "" {
		err = errors.New("未设置数据目录")
		return err
	}

	playersFolderPath, err := utils.FindPlayersFolder(serverSetting.SourceDir)
	if err != nil {
		err = errors.New("无玩家存档数据目录")
		return err
	}

	// 获取目录下的全部文件名及创建日期
	backUpFiles, err := utils.GetSaveFiles(playersFolderPath)
	if err != nil {
		return err
	}

	res := resp.JsonResp(backUpFiles)
	return c.Status(fiber.StatusOK).JSON(res)
}
