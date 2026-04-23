package user

import (
	"go-learn/internal/common/response"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	response.Ok(c, "用户创建成功", nil)
}
