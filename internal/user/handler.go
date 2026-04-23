package user

import (
	"go-learn/internal/common/response"
	"go-learn/internal/common/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv *UserService
}

func NewUserHandler(srv *UserService) *UserHandler {
	return &UserHandler{srv: srv}
}

func (u *UserHandler) Register(c *gin.Context) {
	params, err := validation.BindJSON[UserRegisterReq](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	user, err := u.srv.Register(params)
	if err != nil {
		response.Fail(c, "注册失败")
		return
	}
	response.OkWithData(c, UserRegisterResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: params.Email,
	})
}
