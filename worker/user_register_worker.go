package worker

import (
	"context"
	"frame/models/rdb"
	"frame/pkg/logger"
	"time"
)

type UserRegisterWorker struct {
	workerChs  [workerNum]chan int64
	ctx        context.Context
	cancelFunc context.CancelFunc
	stop       bool
}

func NewUserRegisterWorker() *UserRegisterWorker {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w := &UserRegisterWorker{
		workerChs:  [workerNum]chan int64{},
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}

	for i := 0; i < workerNum; i++ {
		w.workerChs[i] = make(chan int64, 8)

		go func(idx int) {
			for {
				select {
				case id := <-w.workerChs[idx]:
					logger.Sugar.Info(id)
					// handle some things
				}
			}
		}(i)
	}

	return w
}

func (w *UserRegisterWorker) Start() {
	logger.Logger.Info("UserRegisterWorker started")
	w.runMqListener()
}

func (w *UserRegisterWorker) Stop() {
	logger.Logger.Info("UserRegisterWorker stopping")

	w.stop = true
	w.cancelFunc()

	for i := 0; i < workerNum; i++ {
		for {
			if len(w.workerChs[i]) == 0 {
				break
			}
		}
	}

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

		w.workerChs[id%workerNum] <- id
	}
}
