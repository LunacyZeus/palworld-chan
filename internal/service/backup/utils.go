package backup

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ZipDirectory 将指定目录下的所有文件和文件夹打包成zip文件
func ZipDirectory(sourceDir, destinationFile string) error {
	zipFile, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 将文件路径转换为相对路径
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 如果是目录，只写入头信息
		if info.IsDir() {
			header.Name += "/"
			_, err := zipWriter.CreateHeader(header)
			return err
		}

		// 创建文件
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 打开原文件
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 将原文件内容拷贝到zip文件中
		_, err = io.Copy(writer, file)
		return err
	})
}

func GetNowTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")
	return formattedTime
}

// cleanupOldBackups 删除多余的历史备份文件
func cleanupOldBackups(destinationDir, backupFileName string, maxBackupCount int) {
	backupFiles, err := filepath.Glob(filepath.Join(destinationDir, backupFileName+"*.zip"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 按备份文件名中的时间戳进行排序
	sortBackupFiles(backupFiles)

	// 保留最新的 maxBackupCount 个备份文件，删除多余的
	if len(backupFiles) > maxBackupCount {
		filesToDelete := backupFiles[:len(backupFiles)-maxBackupCount]
		for _, fileToDelete := range filesToDelete {
			err := os.Remove(fileToDelete)
			if err != nil {
				fmt.Println("Error deleting old backup:", err)
			} else {
				//fmt.Println("Old backup deleted:", fileToDelete)
			}
		}
	}
}

// sortBackupFiles 按备份文件名中的时间戳进行排序
func sortBackupFiles(files []string) {
	sort.Slice(files, func(i, j int) bool {
		return extractTimestamp(files[i]) < extractTimestamp(files[j])
	})
}

// extractTimestamp 从备份文件名中提取时间戳
func extractTimestamp(filename string) int64 {
	// 文件名格式为 backup_20060102_150405.zip
	parts := strings.Split(filename, "_")
	if len(parts) != 3 {
		return 0
	}

	timestampStr := parts[1]
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0
	}

	return timestamp
}
