package middleware

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/email"
	"code.coolops.cn/blog_services/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 自定义捕获异常Recovery
func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recovery err: %v"
				global.Logger.WithCallerFrames().ErrorF(s, err)
				err := defailtMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.ErrorF("mail.SendEmail err: %v", err)
				}
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
