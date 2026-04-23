package router

import (
	"go-learn/internal/common/storage"
	"go-learn/internal/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup, infra *storage.Infra) {
	user.RegisterRouter(r, infra)
}
