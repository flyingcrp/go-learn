package router

import (
	"go-learn/app/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	user.RegisterRouter(r)
}
