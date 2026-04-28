package role

import (
	"go-learn/internal/common/response"
	"go-learn/internal/common/validation"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	srv *RoleService
}

func NewRoleHandler(srv *RoleService) *RoleHandler {
	return &RoleHandler{srv: srv}
}

func (h *RoleHandler) Create(c *gin.Context) {
	params, err := validation.BindJSON[RoleCreateRequest](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	data, err := h.srv.Create(params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, data)
}

func (h *RoleHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	data, err := h.srv.Detail(id)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	if data == nil {
		response.OkWithData(c, nil)
		return
	}
	response.OkWithData(c, toDetailResponse(data))
}
