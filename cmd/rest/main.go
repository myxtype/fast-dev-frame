package main

import (
	"flag"
	"frame/rest"
)

func main() {
	flag.Parse()

	rest.StartServer()
}
