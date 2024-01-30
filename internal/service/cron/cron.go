package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

var once sync.Once

func StartCron() {
	once.Do(func() {
		c := cron.New()
		// 每5秒执行一次
		c.AddFunc("@every 5s", func() {
			fmt.Printf("Every one seconds, %s\n", time.Now().Format("15:04:05"))
		})
		c.Start()
	})
}
