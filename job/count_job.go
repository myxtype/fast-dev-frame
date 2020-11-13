package job

import (
	"frame/pkg/logger"
	"time"
)

type CountJob struct {
	running bool // 需要自己控制任务是否重复
}

func (job *CountJob) Run() {
	if job.running {
		return
	}
	job.running = true
	defer func() {
		job.running = false
	}()

	logger.Logger.Info("count job in")
	time.Sleep(6 * time.Second)
	logger.Logger.Info("count job out")
}
