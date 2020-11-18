package worker

import (
	"frame/models/rdb"
	"frame/pkg/logger"
	"frame/pkg/queworker"
	"frame/pkg/worker"
	"os"
	"os/signal"
	"syscall"
)

func StartWorker() {
	wm := worker.NewWorkerManager()

	// 用户注册成功后的消息
	wm.AddWorker(queworker.NewQueueWorker(rdb.Shared().NewQueue("registered"), &UserRegisterHandler{}))
	// to add more here

	wm.Start()
	logger.Logger.Info("All worker started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down worker...")

	wm.Stop()
	wm.Wait()
}
