package backup

import (
	"fmt"
	"os"
	"palworld-chan/pkg/logger"
)

type BackupService struct {
	name   string
	token  string
	nodeID string
}

func New(name string, token string, nodeID string) BackupService {
	return BackupService{
		name:   name,
		token:  token,
		nodeID: nodeID,
	}
}

func LocalBackUp(sourceDir, destinationDir string, maxBackupCount int) {
	//sourceDir := "/path/to/source/directory" // 源目录
	//destinationDir := "/path/to/backup"      // 备份文件位置 不带文件名

	//maxBackupCount := 5      // 最大历史备份数量
	backupFileName := "pal_" // 备份文件名前缀

	// 检测源目录是否存在
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		logger.Fatal("Error: Source directory does not exist.")
		return
	}

	// 检测目标目录是否存在
	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		logger.Fatal("Error: Destination directory does not exist.")
		return
	}

	timeStr := GetNowTime()
	destinationFilePath := fmt.Sprintf("%s/%s%s.zip", destinationDir, backupFileName, timeStr)

	logger.Info("Backup file now in %s", destinationFilePath)

	err := ZipDirectory(sourceDir, destinationFilePath)
	if err != nil {
		logger.Error("Error:%v", err)
		return
	}

	logger.Info("Backup successful!")

	// 删除多余的历史备份文件
	cleanupOldBackups(destinationDir, backupFileName, maxBackupCount)
}
