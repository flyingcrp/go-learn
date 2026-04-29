package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TraceIdKey struct{}

func TraceGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := uuid.NewV7()
		uuidStr := uid.String()
		newCtx := context.WithValue(c.Request.Context(), TraceIdKey{}, uuidStr)
		c.Request = c.Request.WithContext(newCtx)
		c.Writer.Header().Set("X-Trace-Id", uuidStr)
		c.Next()
	}
}

func MustGetTraceId(c *gin.Context) string {
	return c.MustGet(TraceIdKey{}).(string)
}
