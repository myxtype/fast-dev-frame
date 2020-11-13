package main

import (
	"flag"
	"frame/job"
)

func main() {
	flag.Parse()

	job.StartJob()
}
