package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("1704179023654_1356"), nil
		}, jwt.WithValidMethods([]string{"HS256"}))
		if err != nil || !jwtToken.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		c.Next()
	}
}
