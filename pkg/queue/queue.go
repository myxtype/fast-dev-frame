package queue

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
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

// Push 向队列中添加任务
func (q *Queue) Push(ctx context.Context, v interface{}) error {
	msg, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return q.PushJob(ctx, msg)
}

// Pop 取出一个任务
func (q *Queue) Pop(ctx context.Context, timeout time.Duration) ([]byte, error) {
	result, err := q.db.BRPop(ctx, timeout, q.formatKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	msg := []byte(result[1])

	return msg, nil
}

// PushJob 发布任务
func (q *Queue) PushJob(ctx context.Context, msg []byte) error {
	return q.db.LPush(ctx, q.formatKey, msg).Err()
}
