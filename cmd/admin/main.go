package main

import (
	"flag"
	"frame/admin"
)

func main() {
	flag.Parse()

	admin.StartServer()
}
