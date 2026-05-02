package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type traceIdKey struct{}

func TraceGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("X-Trace-Id")
		if traceId == "" {
			uid, err := uuid.NewV7()
			if err != nil {
				c.Next()
				return
			}
			traceId = uid.String()
		}
		c.Set(traceIdKey{}, traceId)
		newCtx := context.WithValue(c.Request.Context(), traceIdKey{}, traceId)
		c.Request = c.Request.WithContext(newCtx)
		c.Writer.Header().Set("X-Trace-Id", traceId)
		c.Next()
	}
}

func GetTraceId(ctx context.Context) string {
	if v := ctx.Value(traceIdKey{}); v == nil {
		return ""
	} else {
		return v.(string)
	}
}
