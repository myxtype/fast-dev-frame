package redisdb

import (
	"context"
	"frame/conf"
	"frame/pkg/queue"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	store *Store
	once  sync.Once
)

type Store struct {
	client *redis.Client
}

// 单例模式
func Shared() *Store {
	once.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
	})
	return store
}

func initDb() error {
	cfg := conf.GetConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	store = NewStore(client)

	return nil
}

func NewStore(c *redis.Client) *Store {
	return &Store{client: c}
}

// 获取Redis客户端
func (s *Store) DB() *redis.Client {
	return s.client
}

// Redis健康检查
func (s *Store) Ping() error {
	return s.client.Ping(context.Background()).Err()
}

// 获取分布式锁对象
func (s *Store) Locker() *redislock.Client {
	return redislock.New(s.client)
}

// 获取同步队列对象
func (s *Store) NewQueue(name string) *queue.Queue {
	return queue.NewQueue(name, s.client)
}

// 获取延迟队列对象
func (s *Store) NewDelayQueue(name string) *queue.DelayQueue {
	return queue.NewDelayQueue(name, s.client)
}
