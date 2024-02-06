package tasks

import (
	"palworld-chan/internal/service/dao"
	"palworld-chan/pkg/logger"
	"sync"
	"time"
)

var onceOnline sync.Once
var lastOnlineTimeInstance time.Time

func LastOnlineTime() time.Time {
	onceOnline.Do(func() {
		lastOnlineTimeInstance = time.Now()
	})
	return lastOnlineTimeInstance
}

func GetOnlineUser() (err error) { //监控服务端
	users, err := dao.GetUsersFromRcon()
	if err != nil {
		return
	}
	now := time.Now()
	interval := now.Sub(LastOnlineTime())
	if int64(interval.Seconds()) >= 60 {
		logger.Info("获取到实时在线人数: %d人", len(users))
		lastOnlineTimeInstance = time.Now() //重置时间
	}
	return
}
