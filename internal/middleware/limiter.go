package middleware

import (
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/errcode"
	"code.coolops.cn/blog_services/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// 限流中间件

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
