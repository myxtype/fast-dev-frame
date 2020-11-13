package rest

import "frame/conf"

func StartServer() {
	cfg := conf.GetConfig()

	httpServer := NewHttpServer(cfg.RestServer.Addr)
	httpServer.Start()
}
