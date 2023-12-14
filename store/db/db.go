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

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
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
	cfg := conf.Get().DataSource

	c := &gorm.Config{}
	if cfg.LogDisabled {
		c.Logger = logger.Discard
	}

	// 打开数据库
	gdb, err := gorm.Open(postgres.New(postgres.Config{DSN: cfg.DSN}), c)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	// 设置空闲连接池中连接的最大数量
	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	}
	// 设置打开数据库连接的最大数量
	if cfg.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	}
	// 设置连接可复用的最大时间
	if cfg.MaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.MaxLifetime * time.Second)
	}

	// 迁移
	if cfg.Migrate {
		gdb.AutoMigrate(&model.AdminLog{})
		gdb.AutoMigrate(&model.AdminUser{})
		gdb.AutoMigrate(&model.AdminRole{})

		gdb.AutoMigrate(&model.User{})
	}

	return NewStore(gdb), nil
}

// BeginTx 开启事务
func (s *Store) BeginTx() (*Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

// Rollback 回滚事务
func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

// Commit 提交
func (s *Store) Commit() error {
	return s.db.Commit().Error
}

// Ping 数据库健康检查
func (s *Store) Ping() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
