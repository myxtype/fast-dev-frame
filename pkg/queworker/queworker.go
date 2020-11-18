package queworker

import (
	"context"
	"frame/pkg/logger"
	"frame/pkg/queue"
	"sync"
	"time"
)

type QueueWorkerHandler interface {
	Handle(job *queue.QueueJob)
	Wait()
}

type QueueWorker struct {
	workerCh   chan *queue.QueueJob
	que        *queue.Queue
	ctx        context.Context
	cancelFunc context.CancelFunc
	stop       bool
	wg         sync.WaitGroup
	handler    QueueWorkerHandler
}

type QueueWorkerConfig struct {
	Buffer int
}

func DefaultQueueWorkerConfig() *QueueWorkerConfig {
	return &QueueWorkerConfig{Buffer: 10}
}

func NewQueueWorker(que *queue.Queue, handler QueueWorkerHandler, conf ...*QueueWorkerConfig) *QueueWorker {
	ctx, cancelFunc := context.WithCancel(context.Background())

	var c *QueueWorkerConfig
	if len(conf) > 0 {
		c = conf[0]
	} else {
		c = DefaultQueueWorkerConfig()
	}

	w := &QueueWorker{
		workerCh:   make(chan *queue.QueueJob, c.Buffer),
		que:        que,
		ctx:        ctx,
		cancelFunc: cancelFunc,
		stop:       false,
		wg:         sync.WaitGroup{},
		handler:    handler,
	}

	go func() {
		for {
			select {
			case job := <-w.workerCh:
				w.Do(job)
			}
		}
	}()

	return w
}

func (w *QueueWorker) Do(job *queue.QueueJob) {
	defer func() {
		w.wg.Done()
	}()
	w.handler.Handle(job)
}

func (w *QueueWorker) Start() {
	w.runMqListener()
}

func (w *QueueWorker) Stop() {
	w.stop = true
	w.cancelFunc()

	w.wg.Wait()
}

func (w *QueueWorker) runMqListener() {
	for {
		if w.stop {
			return
		}
		job, err := w.que.Pop(w.ctx, 5*time.Second)
		if job == nil {
			if err != nil {
				logger.Sugar.Error(err)
				time.Sleep(1 * time.Second) // wait 2 second
			}
			continue
		}

		w.wg.Add(1)
		w.workerCh <- job
	}
}
