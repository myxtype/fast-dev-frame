package main

import (
	"frame/admin"
	"frame/conf"
)

func main() {
	httpServer := admin.NewHttpServer(conf.Get().AdminServer.Addr)
	httpServer.Start()
}
