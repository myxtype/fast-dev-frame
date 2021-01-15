package worker

import (
	"frame/pkg/logger"
	"sync"
)

type WorkerManager struct {
	sw          *sync.WaitGroup
	WorkerSlice []Worker
}

func NewWorkerManager() *WorkerManager {
	workerManager := WorkerManager{
		sw: &sync.WaitGroup{},
	}
	workerManager.WorkerSlice = make([]Worker, 0, 10)
	return &workerManager
}

// 添加Worker
func (m *WorkerManager) AddWorker(w Worker) {
	m.WorkerSlice = append(m.WorkerSlice, w)
}

// 全部开始
func (m *WorkerManager) Start() {
	m.sw.Add(len(m.WorkerSlice))
	for _, worker := range m.WorkerSlice {
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

// 全部结束
func (m *WorkerManager) Stop() {
	for _, worker := range m.WorkerSlice {
		go func(w Worker) {
			defer func() {
				m.sw.Done()
			}()
			w.Stop()
		}(worker)
	}
}

// 等待全部结束
func (m *WorkerManager) Wait() {
	m.sw.Wait()
}
