package grace

import (
	"context"
	"frame/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 启动http服务，关闭重启时优雅的关闭
func HttpRun(addr string, handler http.Handler, timeout time.Duration) {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		logger.Sugar.Infof("listen: %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down server...")

	// The context is used to inform the server it has `timeout` to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar.Fatal("Server forced to shutdown:", err)
	}

	logger.Logger.Info("Shut down server ok")
}
