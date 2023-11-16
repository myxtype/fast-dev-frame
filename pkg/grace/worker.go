package grace

import (
	"frame/pkg/logger"
	"frame/pkg/worker"
	"os"
	"os/signal"
	"syscall"
)

// WorkerRun 优雅地关闭
func WorkerRun(m *worker.WorkerManager) {
	m.Start()
	logger.Logger.Info("All worker started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down worker...")

	m.Stop()
	m.Wait()

	logger.Logger.Info("Shut down worker ok")
}
