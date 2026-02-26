package storage

import (
	"os"
	"time"

	color "github.com/fatih/color"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
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
	dbIns, err := gorm.Open(gormMysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败: " + err.Error())
	}
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
