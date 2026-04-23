package user

import (
	"go-learn/internal/common/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv *UserService
}

func NewUserHandler(srv *UserService) *UserHandler {
	return &UserHandler{srv: srv}
}
func (u *UserHandler) Register(c *gin.Context) {
	u.srv.Register()
	response.Ok(c, "用户创建成功", nil)
}
