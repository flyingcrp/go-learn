package config

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	MySQL           mysql.Config
	JWTSecret       string
}

func Load() (*Config, error) {
	d, _ := time.ParseDuration(os.Getenv("SHUTDOWN_TIMEOUT"))
	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_PASS")
	mysqlAddr := os.Getenv("DB_ADDR")
	if mysqlAddr == "" || mysqlUser == "" || mysqlPwd == "" || os.Getenv("DB_NAME") == "" {
		return nil, fmt.Errorf("mysql 配置异常")
	}
	return &Config{
		Port:            ":9000",
		ShutdownTimeout: d,
		JWTSecret:       os.Getenv("JWT_SECRET"),
		MySQL: mysql.Config{
			User:                 mysqlUser,
			Passwd:               mysqlPwd,
			Addr:                 mysqlAddr,
			Net:                  "tcp",
			ParseTime:            true,
			DBName:               os.Getenv("DB_NAME"),
			Loc:                  time.Local,
			AllowNativePasswords: true},
	}, nil
}
