package storage

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

func NewMySQL(cfg mysql.Config) (*gorm.DB, error) {
	var dbIns *gorm.DB
	var err error
	if cfg.Addr == "" || cfg.User == "" || cfg.Passwd == "" || cfg.DBName == "" {
		err = fmt.Errorf("database config is invalid")
		return nil, err
	}
	dbIns, err = gorm.Open(gm.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("GORM初始化失败: %w", err)
		return nil, err
	}
	sqlDB, err := dbIns.DB()
	if err != nil {
		err = fmt.Errorf("获取SQL连接失败: %v", err)
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		err = fmt.Errorf("数据库连接验证失败: %v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	var version string
	dbIns.Raw("select version()").Scan(&version)
	slog.Info("[MySQL] 链接成功!", "DB", cfg.DBName, "version", version)
	return dbIns, nil
}

func CloseMySQL(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	slog.Info("[MySQL] 已正常关闭")
	return sqlDB.Close()
}
