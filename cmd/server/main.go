package main

import (
	"context"
	"go-learn/internal/common/validation"
	"go-learn/internal/config"
	"go-learn/internal/department"
	"go-learn/internal/middleware"
	"go-learn/internal/role"
	"go-learn/internal/storage"
	"go-learn/internal/user"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	validation.InitTrans()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	infra, cleanup, err := storage.NewInfra(storage.Config{
		MySQL: cfg.MySQL,
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
	v1 := routers.Group("/v1", middleware.TraceGuard())
	{

		authGuard := middleware.AuthGuard(cfg.JWTSecret)
		//注入部门模块
		dep := department.NewDepartmentModule(infra)
		department.RegisterRouter(v1, dep.Handler, authGuard)

		roleModule := role.NewRoleModule(infra)
		role.RegisterRouter(v1, roleModule.Handler, authGuard)

		// 注入用户模块
		userHandler := user.NewUserModule(infra, dep.Utils, roleModule.Utils, cfg.JWTSecret)
		user.RegisterRouter(v1, userHandler, authGuard)
	}

	srv := &http.Server{
		Addr:    cfg.Port,
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
	if cfg.ShutdownTimeout != 0 {
		timeout = cfg.ShutdownTimeout
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		color.Red("服务强制关闭: ", err)
	}
	color.Yellow("服务已退出了")
}
