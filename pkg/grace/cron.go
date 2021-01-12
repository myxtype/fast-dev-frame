package grace

import (
	"frame/pkg/logger"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Job 优雅的关闭
func CronRun(c *cron.Cron, timeout time.Duration) {
	c.Start()
	logger.Logger.Info("Cron started")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down cron...")

	ctx := c.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Logger.Info("Shut down cron success.")
			return
		case <-time.NewTicker(timeout).C:
			logger.Logger.Error("Shutting down cron timeout")
			return
		}
	}
}
