package mysql

import (
	"fmt"
	"frame/conf"
	"frame/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var store *Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
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

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := conf.GetConfig()

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.DataSource.User, cfg.DataSource.Password, cfg.DataSource.Addr, cfg.DataSource.Database)

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}

	// 设置空闲连接池中连接的最大数量
	if cfg.DataSource.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.DataSource.MaxIdle)
	}
	// 设置打开数据库连接的最大数量
	if cfg.DataSource.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(cfg.DataSource.MaxOpen)
	}
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 迁移
	if cfg.DataSource.Migrate {
		gdb.AutoMigrate(&models.User{})
		gdb.AutoMigrate(&models.AdminUser{})
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
