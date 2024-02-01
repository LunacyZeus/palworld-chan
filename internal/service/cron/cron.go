package cron

import (
	"github.com/robfig/cron/v3"
	"palworld-chan/internal/service/cron/job"
	"sync"
)

var once sync.Once

func StartCron() {
	once.Do(func() {
		c := cron.New()
		// 每5秒执行一次
		c.AddFunc("@every 5s", func() {
			_ = job.AutoBackUp()
		})
		// 每10秒执行一次 系统性能检测
		c.AddFunc("@every 10s", func() {
			_ = job.MonitorServer()
		})
		c.Start()
	})
}
