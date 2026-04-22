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

	"go-learn/internal/common/middleware"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	validation.InitTrans()
	storage.Init()
	defer storage.Close()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    500,
			"message": "服务器错误",
			"data":    nil,
		})
	}))
	v1 := r.Group("/v1")
	{
		v1.Use(middleware.AuthGuard())
		router.InitRouter(v1)
	}
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
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
	timeout := 5 * time.Second
	if t := os.Getenv("SHUTDOWN_TIMEOUT"); t != "" {
		d, _ := time.ParseDuration(t)
		if d > 0 {
			timeout = d
		}
	}

	color.Yellow("正在关闭服务...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		color.Red("服务强制关闭: ", err)
	}
	color.Green("服务已退出了")
}
