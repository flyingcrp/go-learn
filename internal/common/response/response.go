package response

import (
	"go-learn/internal/common/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok 成功响应
func Ok(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": message,
		"data":    data,
	})
}
func OkWithData(c *gin.Context, data any) {
	Ok(c, "操作成功", data)
}
func OkWithMessage(c *gin.Context, message string) {
	Ok(c, message, nil)
}

// Fail 失败响应
func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"data":    nil,
	})
}

func FailWithValid(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, validation.Translate(err))
}
