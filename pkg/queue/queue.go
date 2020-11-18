package queue

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type Queue struct {
	db        *redis.Client
	jobName   string
	formatKey string
}

func NewQueue(name string, db *redis.Client) *Queue {
	return &Queue{
		db:        db,
		jobName:   name,
		formatKey: "queue:" + name,
	}
}

// 向队列中添加任务
func (q *Queue) Push(msg interface{}) error {
	job, err := NewDelayQueueJob(msg)
	if err != nil {
		return err
	}

	return q.PushJob(job)
}

// 取出一个任务
func (q *Queue) Pop(ctx context.Context, timeout time.Duration) (*QueueJob, error) {
	result, err := q.db.BRPop(ctx, timeout, q.formatKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var job QueueJob
	if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
		return nil, err
	}

	return &job, nil
}

func (q *Queue) PushJob(job *QueueJob) error {
	return q.db.LPush(context.Background(), q.formatKey, job.Bytes()).Err()
}
