package main

import (
	"flag"
	"frame/worker"
)

func main() {
	flag.Parse()

	worker.StartWorker()
}
