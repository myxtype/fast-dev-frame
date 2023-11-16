package db

import (
	"frame/conf"
	"frame/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	store *Store
	once  sync.Once
)

type Store struct {
	db *gorm.DB
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

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := conf.GetConfig().DataSource

	c := &gorm.Config{}
	if cfg.LogDisabled {
		c.Logger = logger.Discard
	}

	// 打开数据库
	gdb, err := gorm.Open(postgres.New(postgres.Config{DSN: cfg.DSN}), c)
	if err != nil {
		return err
	}

	sdb, err := gdb.DB()
	if err != nil {
		return err
	}

	// 设置空闲连接池中连接的最大数量
	if cfg.MaxIdle > 0 {
		sdb.SetMaxIdleConns(cfg.MaxIdle)
	}
	// 设置打开数据库连接的最大数量
	if cfg.MaxOpen > 0 {
		sdb.SetMaxOpenConns(cfg.MaxOpen)
	}
	// 设置连接可复用的最大时间
	if cfg.MaxLifetime > 0 {
		sdb.SetConnMaxLifetime(cfg.MaxLifetime * time.Second)
	}

	// 迁移
	if cfg.Migrate {
		gdb.AutoMigrate(&model.AdminUser{})
		gdb.AutoMigrate(&model.User{})
	}

	store = NewStore(gdb)

	return nil
}

// 开启事务
func (s *Store) BeginTx() (*Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

// 回滚事务
func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

// 提交
func (s *Store) Commit() error {
	return s.db.Commit().Error
}

// 数据库健康检查
func (s *Store) Ping() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
