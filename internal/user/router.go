package user

import (
	"go-learn/internal/common/middleware"
	"go-learn/internal/common/storage"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, infra *storage.Infra) {
	user := r.Group("/user").Use(middleware.AuthGuard())
	{
		user.POST("/register", Register)
	}

}
