package api

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/service"
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(ctx *gin.Context) {
	param := service.AuthRequest{
		AppKey: ctx.GetHeader("app_key"),
		AppSecret: ctx.GetHeader("app_secret"),
	}
	response := app.Response{Ctx: ctx}
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(ctx)
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.ErrorF("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.ErrorF("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
