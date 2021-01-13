package worker

import (
	"frame/pkg/grace"
	"frame/pkg/queworker"
	"frame/pkg/worker"
	"frame/store/redisdb"
)

func StartWorker() {
	wm := worker.NewWorkerManager()

	// 用户注册成功后的消息
	wm.AddWorker(queworker.NewQueueWorker(redisdb.Shared().NewQueue("registered"), &UserRegisterHandler{}))
	// To add more here

	grace.WorkerRun(wm)
}
