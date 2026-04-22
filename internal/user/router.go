package user

import (
	"go-learn/internal/common/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	user := r.Group("/user").Use(middleware.AuthGuard())
	{
		user.GET("", List)
		user.POST("/register", Register)
		user.GET("/:id", Detail)
		user.POST("/:id/update", Update)
	}

}
