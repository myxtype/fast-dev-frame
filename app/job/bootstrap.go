package job

import (
	"frame/pkg/grace"
	"github.com/robfig/cron/v3"
	"time"
)

// 启动所有定时任务
func StartJob() {
	c := cron.New(cron.WithSeconds())

	c.Schedule(cron.Every(5*time.Second), &CountJob{})

	// 优雅的启动任务
	// 如果当前有任务正在执行，给它30秒的时间
	grace.CronRun(c, 30*time.Second)
}
