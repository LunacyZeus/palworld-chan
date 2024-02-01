package job

import (
	"github.com/duke-git/lancet/v2/convertor"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/backup"
	"palworld-chan/internal/service/dao"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/utils"
	"sync"
	"time"
)

var lastBackUpTimeInstance time.Time
var onceBackUp sync.Once

func LastBackUpTime() time.Time {
	onceBackUp.Do(func() {
		lastBackUpTimeInstance = time.Now()
	})
	return lastBackUpTimeInstance
}

func AutoBackUp() (err error) { //自动备份
	//从持久层解析到结构体
	serverSettingJson, err := dao.Get(consts.BUCKET, "serverSetting")

	if err != nil {
		return
	}

	// 创建 UpdateServerSettingInput 实例
	var updateServerSettingInput models.UpdateServerSettingInput

	// 解析 JSON 字符串到结构体
	err = utils.FromJSONString(serverSettingJson, &updateServerSettingInput)
	if err != nil {
		return
	}

	BackupTimeStr := updateServerSettingInput.BackupTime
	BackupTime, err := convertor.ToInt(BackupTimeStr)
	if err != nil {
		return
	}

	now := time.Now()
	interval := now.Sub(LastBackUpTime())

	if int64(interval.Seconds()) >= BackupTime {
		SourceDir := updateServerSettingInput.SourceDir
		DestDir := updateServerSettingInput.DestDir
		BackupCountStr := updateServerSettingInput.BackupCount

		AccessToken := updateServerSettingInput.AccessToken
		SecretKey := updateServerSettingInput.SecretKey
		Bucket := updateServerSettingInput.Bucket

		if SourceDir != "" && DestDir != "" {
			BackupCount, err := convertor.ToInt(BackupCountStr)
			if err != nil {
				BackupCount = 200
			}
			lastBackUpTimeInstance = time.Now() //重置时间
			logger.Debug("开始本地备份,源目录(%s),备份目录(%s),最大保留数(%d)", SourceDir, DestDir, BackupCount)
			err = backup.LocalBackUp(SourceDir, DestDir, int(BackupCount))
			if err != nil {
				logger.Error("本地备份异常: %v", err)
			}
		}

		if AccessToken != "" && SecretKey != "" && Bucket != "" {

		}

	}
	return
}
