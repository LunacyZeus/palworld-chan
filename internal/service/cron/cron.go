package cron

import (
	"github.com/robfig/cron/v3"
	"palworld-chan/internal/service/cron/tasks"
	"sync"
)

var once sync.Once

func StartCron() {
	once.Do(func() {
		c := cron.New()
		// 每5秒执行一次
		c.AddFunc("@every 5s", func() {
			_ = tasks.AutoBackUp()
		})
		// 每10秒执行一次 系统性能检测
		c.AddFunc("@every 10s", func() {
			_ = tasks.MonitorServer()
		})
		// 每60秒执行一次 获取在线玩家列表
		c.AddFunc("@every 60s", func() {
			_ = tasks.GetOnlineUser()
		})
		c.Start()
	})
}
