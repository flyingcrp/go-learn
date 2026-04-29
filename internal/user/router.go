package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, h *UserHandler, authGuard gin.HandlerFunc) {
	user := r.Group("/user")
	{
		user.POST("/register", h.Register)
		user.POST("/:id/update", authGuard, h.Update)
		user.GET("/:id", authGuard, h.Detail)
		user.POST("/login", h.Login)
	}
}
