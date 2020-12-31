package worker

import (
	"frame/pkg/logger"
	"sync"
)

type WorkerManager struct {
	sync.WaitGroup
	WorkerSlice []Worker
}

func NewWorkerManager() *WorkerManager {
	workerManager := WorkerManager{}
	workerManager.WorkerSlice = make([]Worker, 0, 10)
	return &workerManager
}

func (wm *WorkerManager) AddWorker(w Worker) {
	wm.WorkerSlice = append(wm.WorkerSlice, w)
}

func (wm *WorkerManager) Start() {
	wm.Add(len(wm.WorkerSlice))
	for _, worker := range wm.WorkerSlice {
		go func(w Worker) {
			defer func() {
				err := recover()
				if err != nil {
					logger.Sugar.Error(err)
				}
			}()
			w.Start()
		}(worker)
	}
}

func (wm *WorkerManager) Stop() {
	for _, worker := range wm.WorkerSlice {
		go func(w Worker) {
			defer func() {
				err := recover()
				if err != nil {
					logger.Sugar.Error(err)
				}
			}()

			w.Stop()
			wm.Done()
		}(worker)
	}
}
