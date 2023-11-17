package main

import (
	"frame/conf"
	"frame/internal/rest"
)

func main() {
	httpServer := rest.NewHttpServer(conf.Get().RestServer.Addr)
	httpServer.Start()
}
