package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, h *UserHandler) {
	user := r.Group("/user")
	{
		user.POST("/register", h.Register)
		user.POST("/:id/update", h.Update)
	}
}
