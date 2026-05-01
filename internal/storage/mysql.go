package storage

import (
	"fmt"
	"time"

	color "github.com/fatih/color"
	"github.com/go-sql-driver/mysql"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(cfg mysql.Config) (*gorm.DB, error) {
	if cfg.Addr == "" || cfg.User == "" || cfg.Passwd == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("database config is invalid")
	}
	dbIns, err := gorm.Open(gm.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("GORM初始化失败: %w", err)
	}
	sqlDB, err := dbIns.DB()
	if err != nil {
		return nil, fmt.Errorf("获取SQL连接失败: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接验证失败: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	var version string
	dbIns.Raw("select version()").Scan(&version)
	color.Yellow("[MySQL] 链接成功! DB: %s,version: %s", cfg.DBName, version)
	return dbIns, nil
}

func CloseMySQL(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	color.Yellow("[MySQL] 已正常关闭")
	return sqlDB.Close()
}
