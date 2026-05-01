package role

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.RouterGroup, h *RoleHandler, auth gin.HandlerFunc) {
	role := g.Group("/role").Use(auth)
	{
		role.POST("/create", h.Create)
		role.GET("/:id", h.Detail)
	}
}
