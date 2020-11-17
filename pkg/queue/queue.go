package queue

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type Queue struct {
	db      *redis.Client
	jobName string
}

func NewQueue(name string, db *redis.Client) *Queue {
	return &Queue{
		db:      db,
		jobName: name,
	}
}

// 向队列中添加任务
func (q *Queue) Push(msg interface{}) error {
	job, err := NewDelayQueueJob(msg)
	if err != nil {
		return err
	}

	return q.db.LPush(context.Background(), q.formatKey(), job.Bytes()).Err()
}

// 取出一个任务
func (q *Queue) Pop(timeout time.Duration) (*QueueJob, error) {
	result, err := q.db.BRPop(context.Background(), timeout, q.formatKey()).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var job QueueJob
	if err := json.Unmarshal([]byte(result[0]), &job); err != nil {
		return nil, err
	}

	return &job, nil
}

func (q *Queue) formatKey() string {
	return "queue:" + q.jobName
}
