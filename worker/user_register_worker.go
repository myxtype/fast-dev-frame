package worker

import (
	"context"
	"frame/models/rdb"
	"frame/pkg/logger"
	"sync"
	"time"
)

type UserRegisterWorker struct {
	workerCh   chan interface{}
	ctx        context.Context
	cancelFunc context.CancelFunc
	stop       bool
	wg         sync.WaitGroup
}

func NewUserRegisterWorker() *UserRegisterWorker {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w := &UserRegisterWorker{
		workerCh:   make(chan interface{}, 8),
		ctx:        ctx,
		cancelFunc: cancelFunc,
		wg:         sync.WaitGroup{},
	}

	go func() {
		for {
			select {
			case id := <-w.workerCh:
				w.Do(id)
			}
		}
	}()

	return w
}

// 执行任务
func (w *UserRegisterWorker) Do(param interface{}) {
	defer w.wg.Done()

	time.Sleep(time.Second)
	logger.Sugar.Info(param)
}

func (w *UserRegisterWorker) Start() {
	logger.Logger.Info("UserRegisterWorker started")
	w.runMqListener()
}

func (w *UserRegisterWorker) Stop() {
	logger.Logger.Info("UserRegisterWorker stopping")

	w.stop = true
	w.cancelFunc()

	w.wg.Wait()

	logger.Logger.Info("UserRegisterWorker stopped")
}

// 监听消息队列通知
func (w *UserRegisterWorker) runMqListener() {
	que := rdb.Shared().NewQueue("registered")

	for {
		if w.stop {
			return
		}
		job, err := que.Pop(w.ctx, 10*time.Second)
		if job == nil {
			if err != nil {
				logger.Sugar.Error(err)
				time.Sleep(2 * time.Second) // wait 2 second
			}
			continue
		}

		var id int64
		if err := job.Unmarshal(&id); err != nil {
			logger.Sugar.Error(err, job)
			time.Sleep(time.Second)
			continue
		}

		w.wg.Add(1)
		w.workerCh <- id
	}
}
