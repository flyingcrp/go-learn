package user

import (
	"go-learn/app/common/response"
	"go-learn/app/common/validation"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	_, err := validation.BindJSON[CreateUserReq](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	c.String(200, "done")
}

func List(c *gin.Context) {
	c.JSON(200, gin.H{"message": "user list"})
}
