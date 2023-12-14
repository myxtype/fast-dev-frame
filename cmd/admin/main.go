package main

import (
	"frame/app/admin"
	"frame/conf"
)

func main() {
	httpServer := admin.NewHttpServer(conf.Get().AdminServer.Addr)
	httpServer.Start()
}
