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

	user, err := u.srv.Register(c.Request.Context(), params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, UserRegisterResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: params.Email,
	})
}
func (u *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, "id is required")
		return
	}
	params, err := validation.BindJSON[UserUpdateReq](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	updatedUser, err := u.srv.Update(c.Request.Context(), id, params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, UserUpdateResp{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	})
}

func (u *UserHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, "id is required")
		return
	}
	user, err := u.srv.FindByID(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	if user == nil {
		response.OkWithData(c, nil)
		return
	}
	response.OkWithData(c, toUserDetailResp(user))
}

func (u *UserHandler) Login(c *gin.Context) {
	params, err := validation.BindJSON[UserLoginReq](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	user, err := u.srv.Login(c.Request.Context(), params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	token, err := u.srv.GenerateJWT(user)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, toLoginResp(user, token))
}
func (h *UserHandler) List(c *gin.Context) {
	params, err := validation.BindQuery[UserListReq](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	result, err := h.srv.List(c.Request.Context(), params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, UserListResp{
		List:  result.List,
		Total: result.Total,
	})
}
