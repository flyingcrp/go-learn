package role

import "github.com/gin-gonic/gin"

func RegisterRouter(g *gin.RouterGroup, h *RoleHandler) {
	role := g.Group("/role")
	{
		role.POST("/create", h.Create)
		role.GET("/:id", h.Detail)
	}
}
