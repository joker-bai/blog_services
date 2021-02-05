package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 统一超时时间配置
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), t)
		defer cancel()

		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
