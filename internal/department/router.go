package department

import "github.com/gin-gonic/gin"

func RegisterRouter(g *gin.RouterGroup, h *DepartmentHandler) {
	department := g.Group("/department")
	{
		department.POST("/create", h.Create)
		department.GET("/:id", h.Detail)
	}
}
