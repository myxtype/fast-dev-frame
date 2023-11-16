package worker

import (
	"frame/pkg/grace"
	"frame/pkg/queworker"
	"frame/pkg/worker"
	"frame/store/rdb"
)

func StartWorker() {
	m := worker.NewWorkerManager()

	// 用户注册成功后的消息
	m.AddWorker(queworker.NewQueueWorker(rdb.Shared().NewQueue("registered"), &UserRegisterHandler{}))
	// To add more here

	grace.WorkerRun(m)
}
