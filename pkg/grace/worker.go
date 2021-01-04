package grace

import (
	"frame/pkg/logger"
	"frame/pkg/worker"
	"os"
	"os/signal"
	"syscall"
)

// Worker 优雅的关闭
func WorkerRun(wm *worker.WorkerManager) {
	wm.Start()
	logger.Logger.Info("All worker started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down worker...")

	wm.Stop()
	wm.Wait()

	logger.Logger.Info("Shut down worker ok")
}
