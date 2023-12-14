package rdb

import (
	"context"
	"frame/conf"
	"frame/pkg/queue"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Store struct {
	client *redis.Client
}

func NewStore(c *redis.Client) *Store {
	return &Store{client: c}
}

// Shared 单例模式
var Shared = sync.OnceValue(func() *Store {
	store, err := initDb()
	if err != nil {
		panic(err)
	}
	return store
})

func initDb() (*Store, error) {
	cfg := conf.Get()

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return NewStore(client), nil
}

// DB 获取Redis客户端
func (s *Store) DB() *redis.Client {
	return s.client
}

// Ping Redis健康检查
func (s *Store) Ping() error {
	return s.client.Ping(context.Background()).Err()
}

// Locker 获取分布式锁对象
func (s *Store) Locker() *redislock.Client {
	return redislock.New(s.client)
}

// NewQueue 获取同步队列对象
func (s *Store) NewQueue(name string) *queue.Queue {
	return queue.NewQueue(name, s.client)
}

// NewDelayQueue 获取延迟队列对象
func (s *Store) NewDelayQueue(name string) *queue.DelayQueue {
	return queue.NewDelayQueue(name, s.client)
}
