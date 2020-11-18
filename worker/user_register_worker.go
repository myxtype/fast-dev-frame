package worker

import (
	"context"
	"frame/models/rdb"
	"frame/pkg/logger"
	"time"
)

type UserRegisterWorker struct {
	workerChs [workerNum]chan int64
}

func NewUserRegisterWorker() *UserRegisterWorker {
	w := &UserRegisterWorker{workerChs: [workerNum]chan int64{}}

	for i := 0; i < workerNum; i++ {
		w.workerChs[i] = make(chan int64, 8)

		go func(idx int) {
			for {
				select {
				case id := <-w.workerChs[idx]:
					logger.Sugar.Info(id)
					// handle do something
				}
			}
		}(i)
	}

	return w
}

func (w *UserRegisterWorker) Start() {
	go w.runMqListener()
	logger.Logger.Info("UserRegisterWorker Start")
}

// 监听消息队列通知
func (w *UserRegisterWorker) runMqListener() {
	q := rdb.Shared().NewQueue("registered")

	for {
		job, err := q.Pop(context.Background(), 10*time.Second)
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
