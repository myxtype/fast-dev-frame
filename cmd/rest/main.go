package main

import (
	"frame/conf"
	"frame/rest"
)

func main() {
	httpServer := rest.NewHttpServer(conf.Get().RestServer.Addr)
	httpServer.Start()
}
