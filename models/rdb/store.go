package rdb

import (
	"context"
	"frame/conf"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"sync"
)

var store *Store
var storeOnce sync.Once

type Store struct {
	client *redis.Client
}

// 单例模式
func Shared() *Store {
	storeOnce.Do(func() {
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

// 获取分布式锁对象
func (s *Store) Locker() *redislock.Client {
	return redislock.New(s.client)
}
