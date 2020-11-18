package worker

import (
	"frame/pkg/logger"
	"frame/pkg/worker"
	"os"
	"os/signal"
	"syscall"
)

func StartWorker() {
	wm := worker.NewWorkerManager()

	wm.AddWorker(NewUserRegisterWorker())
	// add more

	wm.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down worker...")

	wm.Stop()
	wm.Wait()
}
