package middleware

import "github.com/gin-gonic/gin"

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("userID", 123)
		c.Next()
	}
}
