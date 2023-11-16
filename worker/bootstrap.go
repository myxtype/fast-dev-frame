package worker

import (
	"frame/pkg/grace"
	"frame/pkg/worker"
)

func StartWorker() {
	m := worker.NewWorkerManager()

	// todo
	// m.AddWorker()

	grace.WorkerRun(m)
}
