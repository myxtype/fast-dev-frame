package queue

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var ErrRemoveNil = errors.New("DelayQueue job remove nil")

type DelayQueue struct {
	db        *redis.Client
	jobName   string
	formatKey string
}

func NewDelayQueue(name string, db *redis.Client) *DelayQueue {
	return &DelayQueue{
		db:        db,
		jobName:   name,
		formatKey: "delay-queue:" + name,
	}
}

// Push 向队列中添加任务
func (q *DelayQueue) Push(ctx context.Context, v interface{}, delayAt time.Time) error {
	msg, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return q.PushJob(ctx, msg, delayAt)
}

// Remove 删除队列中的一个任务
func (q *DelayQueue) Remove(ctx context.Context, v interface{}) error {
	msg, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return q.RemoveJob(ctx, msg)
}

// Pop 取出一个任务
func (q *DelayQueue) Pop(ctx context.Context) ([]byte, error) {
	res, err := q.db.ZRangeByScore(ctx, q.formatKey, &redis.ZRangeBy{
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

	if err = q.RemoveJob(ctx, msg); err != nil {
		if errors.Is(err, ErrRemoveNil) {
			return nil, nil
		}
		return nil, err
	}

	return msg, nil
}

// PushJob 添加任务
func (q *DelayQueue) PushJob(ctx context.Context, msg []byte, delayAt time.Time) error {
	return q.db.ZAdd(ctx, q.formatKey, redis.Z{
		Score:  float64(delayAt.Unix()),
		Member: msg,
	}).Err()
}

// RemoveJob 删除任务
func (q *DelayQueue) RemoveJob(ctx context.Context, msg []byte) error {
	row, err := q.db.ZRem(ctx, q.formatKey, msg).Result()
	if err != nil {
		return err
	}
	if row == 0 {
		return ErrRemoveNil
	}
	return nil
}
