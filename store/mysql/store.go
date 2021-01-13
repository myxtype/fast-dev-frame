package mysql

import (
	"frame/conf"
	"frame/model"
	"gorm.io/driver/mysql"
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
	gdb, err := gorm.Open(mysql.New(mysql.Config{DSN: cfg.DSN, DefaultStringSize: 256}), c)
	if err != nil {
		return err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return err
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
		gdb.AutoMigrate(&model.User{})
		gdb.AutoMigrate(&model.AdminUser{})
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
