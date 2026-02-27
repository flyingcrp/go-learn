package storage

import (
	"os"
	"time"

	color "github.com/fatih/color"
	"github.com/go-sql-driver/mysql"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Addr:                 os.Getenv("DB_ADDR"),
		Net:                  "tcp",
		ParseTime:            true,
		DBName:               os.Getenv("DB_NAME"),
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}
	dbIns, err := gorm.Open(gm.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("GORM初始化失败: " + err.Error())
	}
	sqlDB, err := dbIns.DB()
	if err != nil {
		panic("获取SQL连接失败: " + err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		panic("数据库连接验证失败: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	db = dbIns
	var version string
	db.Raw("select version()").Scan(&version)
	color.Green("[MySQL] 链接成功! Version: %s", version)
}
func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
	color.Green("[MySQL] 链接已关闭!")
}
