package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type ctxKey struct{}

var userCtxKey = &ctxKey{}

type UserInfo struct {
	UserID string
	Name   string
	Email  string
}

func MustGetUserInfo(c context.Context) UserInfo {
	u, ok := c.Value(userCtxKey).(UserInfo)
	if !ok {
		panic("user info not found")
	}
	return u
}
func AuthGuard(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		token := strings.Split(tokenStr, "Bearer ")[1]
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))
		if err != nil || !jwtToken.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		uid, ok := claims["user_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"data":    nil,
				"code":    401,
				"message": "Unauthorized"})
			return
		}
		ctx := context.WithValue(c.Request.Context(), userCtxKey, UserInfo{
			UserID: uid,
			Name:   claims["name"].(string),
			Email:  claims["email"].(string),
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
