package pushing

import (
	"frame/conf"
	"frame/pkg/logger"
	"github.com/myxtype/go-webreal"
)

func Start() {
	hub := webreal.NewSubscriptionHub()

	// 启动服务
	server := webreal.NewServer(&Handler{}, hub, webreal.DefaultConfig())

	cfg := conf.Get().PushingServer

	logger.Sugar.Infof("Start pushing %v/ws", cfg.Addr)

	if err := server.Run(cfg.Addr, "/ws"); err != nil {
		logger.Sugar.Error(err)
	}
}
