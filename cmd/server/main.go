package main

import (
	"context"
	"go-learn/internal/common/validation"
	"go-learn/internal/config"
	"go-learn/internal/domain/department"
	"go-learn/internal/domain/role"
	"go-learn/internal/domain/user"
	"go-learn/internal/infra/logger"
	"go-learn/internal/infra/middleware"
	"go-learn/internal/infra/storage"
	"log/slog"

	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	validation.InitTrans()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	infra, cleanup, err := storage.NewInfra(storage.Config{
		MySQL: cfg.MySQL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize infrastructure,error:%v", err)
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
	v1 := routers.Group("/v1", middleware.TraceGuard())
	{

		authGuard := middleware.AuthGuard(cfg.JWTSecret)
		//注入部门模块
		dep := department.NewDepartmentModule(infra)
		department.RegisterRouter(v1, dep.Handler, authGuard)

		roleModule := role.NewRoleModule(infra)
		role.RegisterRouter(v1, roleModule.Handler, authGuard)

		// 注入用户模块
		userHandler := user.NewUserModule(infra, dep.Srv, roleModule.Srv, cfg.JWTSecret)
		user.RegisterRouter(v1, userHandler, authGuard)
	}

	srv := &http.Server{
		Addr:         cfg.Port,
		Handler:      routers,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Listen Error: ", "error", err)
		}
	}()
	<-ctx.Done()
	// 优雅关闭：开发环境可设置 SHUTDOWN_TIMEOUT=1s 加快重启
	timeout := 600 * time.Millisecond
	if cfg.ShutdownTimeout != 0 {
		timeout = cfg.ShutdownTimeout
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("服务关闭失败: ", "error", err)
	}
	slog.Info("服务已退出")
}
