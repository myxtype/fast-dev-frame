package worker

import (
	"frame/pkg/logger"
	"frame/pkg/queue"
	"time"
)

type UserRegisterHandler struct{}

// 执行任务
func (handler *UserRegisterHandler) Handle(job *queue.QueueJob) {
	var id int64
	if err := job.Unmarshal(&id); err != nil {
		logger.Sugar.Error(err)
		return
	}

	logger.Sugar.Info(id)
	time.Sleep(time.Second)
}

// 可能需要等待未完成的任务
func (handler *UserRegisterHandler) Wait() {
	// todo
}
