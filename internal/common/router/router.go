package router

import (
	"go-learn/internal/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	user.RegisterRouter(r)
}
