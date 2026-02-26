package response

import (
	"go-learn/app/common/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailWithValid(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": validation.Translate(err),
		"data":    nil,
	})
}
