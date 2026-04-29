package department

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.RouterGroup, h *DepartmentHandler, auth gin.HandlerFunc) {
	department := g.Group("/department").Use(auth)
	{
		department.POST("/create", h.Create)
		department.GET("/:id", h.Detail)
	}
}
