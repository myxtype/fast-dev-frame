package main

import (
	"frame/app/rest"
	"frame/conf"
)

func main() {
	httpServer := rest.NewHttpServer(conf.Get().RestServer.Addr)
	httpServer.Start()
}
