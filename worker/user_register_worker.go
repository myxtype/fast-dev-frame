package worker

import (
	"frame/models/rdb"
	"frame/pkg/logger"
	"frame/pkg/queue"
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
					logger.Sugar.Debug(id)
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
	q := queue.NewDelayQueue("registered", rdb.Shared().DB())
	for {
		job, err := q.Pop()
		if job == nil {
			if err != nil {
				logger.Sugar.Error(err)
			}
			time.Sleep(time.Second)
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
