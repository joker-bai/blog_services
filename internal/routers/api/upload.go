package api

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/service"
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/convert"
	"code.coolops.cn/blog_services/pkg/errcode"
	"code.coolops.cn/blog_services/pkg/upload"
	"github.com/gin-gonic/gin"
)

// 上传文件的路由方法
type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(ctx *gin.Context) {
	response := app.Response{Ctx: ctx}
	file, fileHeader, err := ctx.Request.FormFile("file")
	fileType := convert.StrTo(ctx.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	if fileHeader != nil || fileType <= 0 {
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}
	svc := service.NewService(ctx)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, *fileHeader)
	if err != nil {
		global.Logger.ErrorF("svc.UploadFile err: %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

// 多图片上传
func (u Upload) MultiUploadFile(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	forms, err := ctx.MultipartForm()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	files := forms.File["files"]
	fileType := convert.StrTo(ctx.PostForm("type")).MustInt()
	svc := service.NewService(ctx)
	for _, file := range files {
		fileF, _ := file.Open()
		fileInfo, err := svc.UploadFile(upload.FileType(fileType), fileF, *file)
		if err != nil {
			global.Logger.ErrorF("svc.UploadFile err: %v", err)
			errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		response.ToResponse(gin.H{
			"file_access_url": fileInfo.AccessUrl,
		})
	}
}
