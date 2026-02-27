package main

import (
	"context"
	"go-learn/app/common/router"
	"go-learn/app/common/storage"
	"go-learn/app/common/validation"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
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
		router.InitRouter(v1)
	}
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	// 检测是否在air环境中运行
	isAirMode := os.Getenv("AIR_MODE") == "true"
	
	var ctx context.Context
	var stop context.CancelFunc
	
	if isAirMode {
		// Air模式下使用快速关闭
		ctx, stop = context.WithCancel(context.Background())
		go func() {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan
			stop()
		}()
	} else {
		// 正常模式下使用标准优雅关闭
		ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	}
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen Error: %s\n", err)
		}
	}()

	<-ctx.Done()
	
	if isAirMode {
		color.Yellow("Air模式: 快速重启中...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			color.Red("服务关闭错误: ", err)
		}
		color.Green("Air模式: 服务已准备重启")
	} else {
		color.Yellow("接收到退出信号，正在关闭服务...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			color.Red("服务强制关闭: ", err)
		}
		color.Yellow("服务已安全退出")
	}
}