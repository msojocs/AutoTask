package cron

import (
	"fmt"
	"github.com/msojocs/AutoTask/v1/services/task"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

// Init 初始化crontab
func Init() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))))
	_, err := c.AddFunc("0 0 * * * *", func() { // 每隔一秒执行一次
		unix := time.Now().Unix()
		fmt.Printf("111--start, time=%d\n", unix)
		task.Run()
	})
	if err != nil {
		return
	}
	c.Start()
}
