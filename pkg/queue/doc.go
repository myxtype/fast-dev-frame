package queue

//import (
//	"frame/models/rdb"
//	"frame/pkg/logger"
//	"time"
//)
//
//func main() {
//	queue := NewDelayQueue("registered", rdb.Shared().DB())
//
//	// 插入任务
//	if err := queue.Push(10000, time.Now().Add(10*time.Minute)); err != nil {
//		panic(err)
//	}
//
//	// 拉取任务
//	for {
//		job, err := queue.Pop()
//		if job == nil {
//			if err != nil {
//				logger.Sugar.Error(err)
//			}
//			time.Sleep(time.Second)
//			continue
//		}
//
//		var id int64
//		if err := job.Unmarshal(&id); err != nil {
//			logger.Sugar.Error(err)
//			time.Sleep(5 * time.Second)
//			continue
//		}
//
//		// 将id放入特定的chan中，以便快速拉取下一个任务
//		workerChs[id%10] <- id
//	}
//}
