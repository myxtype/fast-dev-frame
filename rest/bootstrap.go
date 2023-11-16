package rest

import "frame/conf"

func StartServer() {
	httpServer := NewHttpServer(conf.Get().RestServer.Addr)
	httpServer.Start()
}
