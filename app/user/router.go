package user

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/user", List)
	r.POST("/user/register", Register)
}
