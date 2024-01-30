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
			job.AutoBackUp()
		})
		c.Start()
	})
}
