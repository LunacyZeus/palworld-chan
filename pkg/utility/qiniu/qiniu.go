package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"palworld-chan/pkg/logger"
)

// 七牛云配置

func UploadFile(localFilePath, key string, accessKey, secretKey, bucket string) {
	//localFilePath := "/path/to/your/file" // 本地文件路径
	//key := "your_custom_key"              // 在七牛中保存的文件名，可选，如果不设置会自动生成

	zone, err := storage.GetZone(accessKey, bucket)
	if err != nil {
		logger.Error("Error:%v", err)
		return
	}

	// 初始化配置
	cfg := storage.Config{
		Zone: zone, // 可根据自己的需求选择不同的 Zone
	}

	// 创建认证信息
	mac := qbox.NewMac(accessKey, secretKey)

	// 构建上传管理器
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 构建上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	// 上传文件
	err = uploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, nil)
	if err != nil {
		logger.Fatal("Error:%v", err)
		return
	}

	logger.Info("File uploaded successfully!")
	logger.Info("File bucket(%s) key(%s)\n", bucket, ret.Key)
}
