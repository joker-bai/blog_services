package middleware

import (
	"bytes"
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
		}
		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access.log: method: %s, status_code: %d, " + "begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).InfoF(s, ctx.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)

	}
}
