package backup

import (
	"errors"
	"fmt"
	"os"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/qiniu"
)

func QiNiuBackUp(sourceDir, destinationDir string, accessKey, secretKey, bucket string) (err error) {
	//sourceDir := "/path/to/source/directory" // 源目录
	//destinationDir := "/path/to/backup"      // 备份文件位置 不带文件名

	// 检测源目录是否存在
	_, err = os.Stat(sourceDir)
	if os.IsNotExist(err) {
		err = errors.New("Error: Source directory does not exist.")
		return
	}

	_, err = os.Stat(destinationDir)

	// 检测目标目录是否存在
	if os.IsNotExist(err) {
		err = errors.New("Error: Destination directory does not exist.")
		return
	}

	timeStr := GetNowTime()
	destinationFilePath := fmt.Sprintf("%s/pal_%s.zip", destinationDir, timeStr)

	logger.Info("Backup file now in %s", destinationFilePath)

	err = ZipDirectory(sourceDir, destinationFilePath)
	if err != nil {
		logger.Error("Error:%v", err)
		return
	} else {
		logger.Info("Backup successful!")
	}

	key := fmt.Sprintf("pal_%s.zip", timeStr)
	qiniu.UploadFile(destinationFilePath, key, accessKey, secretKey, bucket)
	return
}
