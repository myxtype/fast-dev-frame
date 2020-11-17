package admin

import "frame/conf"

func StartServer() {
	cfg := conf.GetConfig()

	httpServer := NewHttpServer(cfg.AdminServer.Addr)
	httpServer.Start()
}
