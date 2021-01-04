package worker

import (
	"frame/models/redisdb"
	"frame/pkg/grace"
	"frame/pkg/queworker"
	"frame/pkg/worker"
)

func StartWorker() {
	wm := worker.NewWorkerManager()

	// 用户注册成功后的消息
	wm.AddWorker(queworker.NewQueueWorker(redisdb.Shared().NewQueue("registered"), &UserRegisterHandler{}))
	// To add more here

	grace.WorkerRun(wm)
}
