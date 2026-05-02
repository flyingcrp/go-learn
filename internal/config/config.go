package config

import (
	"fmt"
	"go-learn/internal/infra/logger"
	"log/slog"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	MySQL           mysql.Config
	JWTSecret       string
	LogLevel        slog.Level
	LogFormat       logger.Format
}

func Load() (*Config, error) {
	d, _ := time.ParseDuration(os.Getenv("SHUTDOWN_TIMEOUT"))
	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_PASS")
	mysqlAddr := os.Getenv("DB_ADDR")
	if mysqlAddr == "" || mysqlUser == "" || mysqlPwd == "" || os.Getenv("DB_NAME") == "" {
		return nil, fmt.Errorf("mysql 配置异常")
	}
	var logLevel slog.Level
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	var logFormat logger.Format
	switch os.Getenv("LOG_FORMAT") {
	case "text":
		logFormat = logger.Text
	default:
		logFormat = logger.JSON
	}
	return &Config{
		Port:            ":9000",
		ShutdownTimeout: d,
		JWTSecret:       os.Getenv("JWT_SECRET"),
		LogLevel:        logLevel,
		LogFormat:       logFormat,
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
