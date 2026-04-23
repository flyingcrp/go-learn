package response

import (
	"go-learn/internal/common/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok 成功响应
func Ok(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}

func FailWithValid(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": validation.Translate(err),
		"data":    nil,
	})
}
