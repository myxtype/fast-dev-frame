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
}

type QueueWorker struct {
	workerCh   chan *queue.QueueJob
	que        *queue.Queue
	ctx        context.Context
	cancelFunc context.CancelFunc
	stop       bool
	wg         sync.WaitGroup
	handler    QueueWorkerHandler
	conf       *QueueWorkerConfig
}

type QueueWorkerConfig struct {
	Buffer      int           // 队列缓冲区大小
	ReadTimeout time.Duration // 队列读取超时时间
	ReadErrWait time.Duration // 队列中读取错误后等待时间
}

func DefaultQueueWorkerConfig() *QueueWorkerConfig {
	return &QueueWorkerConfig{
		Buffer:      10,
		ReadTimeout: 5 * time.Second,
		ReadErrWait: time.Second,
	}
}

func NewQueueWorker(que *queue.Queue, handler QueueWorkerHandler, optionals ...*QueueWorkerConfig) *QueueWorker {
	ctx, cancelFunc := context.WithCancel(context.Background())

	var conf *QueueWorkerConfig
	if len(optionals) > 0 {
		conf = optionals[0]
	} else {
		conf = DefaultQueueWorkerConfig()
	}

	w := &QueueWorker{
		workerCh:   make(chan *queue.QueueJob, conf.Buffer),
		que:        que,
		ctx:        ctx,
		cancelFunc: cancelFunc,
		stop:       false,
		wg:         sync.WaitGroup{},
		handler:    handler,
		conf:       conf,
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
		job, err := w.que.Pop(w.ctx, w.conf.ReadTimeout)
		if job == nil {
			if err != nil {
				logger.Sugar.Error(err)
				time.Sleep(w.conf.ReadErrWait)
			}
			continue
		}

		w.wg.Add(1)
		w.workerCh <- job
	}
}
