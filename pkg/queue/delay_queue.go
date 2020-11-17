package queue

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

/**
 * 基于Redis的延迟队列
 * 适用于简单的Worker任务
 */
type DelayQueue struct {
	db      *redis.Client
	jobName string
}

func NewDelayQueue(name string, db *redis.Client) *DelayQueue {
	return &DelayQueue{
		db:      db,
		jobName: name,
	}
}

// 向队列中添加任务
func (q *DelayQueue) Push(msg interface{}, delayAt time.Time) error {
	job, err := NewDelayQueueJob(msg)
	if err != nil {
		return err
	}

	return q.db.ZAdd(context.Background(), q.formatKey(), &redis.Z{
		Score:  float64(delayAt.Unix()),
		Member: job.Bytes(),
	}).Err()
}

// 取出一个任务
func (q *DelayQueue) Pop() (*DelayQueueJob, error) {
	res, err := q.db.ZRangeByScore(context.Background(), q.formatKey(), &redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.FormatInt(time.Now().Unix(), 10),
		Offset: 0,
		Count:  1,
	}).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	msg := []byte(res[0])

	var job DelayQueueJob
	if err := json.Unmarshal(msg, &job); err != nil {
		return nil, err
	}

	// 删除这个任务
	if row, err := q.db.ZRem(context.Background(), q.formatKey(), msg).Result(); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, nil
	}

	return &job, nil
}

func (q *DelayQueue) formatKey() string {
	return "delay-queue:" + q.jobName
}
