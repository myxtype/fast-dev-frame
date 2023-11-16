package main

import (
	"frame/conf"
	"frame/internal/admin"
)

func main() {
	httpServer := admin.NewHttpServer(conf.Get().AdminServer.Addr)
	httpServer.Start()
}
