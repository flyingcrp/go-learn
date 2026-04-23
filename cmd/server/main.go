package main

import (
	"context"
	"go-learn/internal/common/router"
	"go-learn/internal/common/storage"
	"go-learn/internal/common/validation"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	validation.InitTrans()
	infra, cleanup, err := storage.NewInfra(storage.Config{
		MySQL: mysql.Config{User: os.Getenv("DB_USER"),
			Passwd:               os.Getenv("DB_PASS"),
			Addr:                 os.Getenv("DB_ADDR"),
			Net:                  "tcp",
			ParseTime:            true,
			DBName:               os.Getenv("DB_NAME"),
			Loc:                  time.Local,
			AllowNativePasswords: true},
	})
	if err != nil {
		log.Fatalf("Failed to initialize infrastructure: %v", err)
	}
	defer cleanup()
	routers := gin.New()
	routers.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    500,
			"message": "服务器错误",
			"data":    nil,
		})
	}))
	v1 := routers.Group("/v1")
	{
		router.InitRouter(v1, infra)
	}
	srv := &http.Server{
		Addr:    ":9000",
		Handler: routers,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen Error: %s\n", err)
		}
	}()
	<-ctx.Done()
	// 优雅关闭：开发环境可设置 SHUTDOWN_TIMEOUT=1s 加快重启
	timeout := 600 * time.Millisecond
	if t := os.Getenv("SHUTDOWN_TIMEOUT"); t != "" {
		d, _ := time.ParseDuration(t)
		if d > 0 {
			timeout = d
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		color.Red("服务强制关闭: ", err)
	}
	color.Yellow("服务已退出了")
}
