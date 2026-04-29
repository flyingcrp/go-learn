package department

import (
	"go-learn/internal/common/response"
	"go-learn/internal/common/validation"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	srv *DepartmentService
}

func NewDepartmentHandler(srv *DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{srv: srv}
}
func (h *DepartmentHandler) Create(c *gin.Context) {
	params, err := validation.BindJSON[DepartmentCreateRequest](c)
	if err != nil {
		response.FailWithValid(c, err)
		return
	}
	ctx := c.Request.Context()
	data, err := h.srv.Create(ctx, params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.OkWithData(c, data)
}

func (h *DepartmentHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	data, err := h.srv.Detail(ctx, id)
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
