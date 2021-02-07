package v1

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/service"
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/convert"
	"code.coolops.cn/blog_services/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取多篇文章
// @Produce json
// @Tags 文章
// @Param title query string false "文章名称" maxlength(100)
// @Param desc query string false "文章描述"
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(ctx *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid err: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 业务逻辑
	svc := service.NewService(ctx)
	articles, err := svc.ArticleList(&param)
	if err != nil {
		global.Logger.ErrorF("svc.ArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorListArticleFail)
		return
	}
	totalRows, err := svc.ArticleCount(&service.ArticleCountRequest{
		Title: param.Title,
		State: param.State,
	})
	if err != nil {
		global.Logger.ErrorF("svc.ArticleCount err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	response.ToResponseList(articles, totalRows)
	return
}

// @Summary 新增文章
// @Produce json
// @Tags 文章
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Param desc body string false "文章描述" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(3)
// @Param cover_image_url body string false "封面图片地址" maxlength(255)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(ctx *gin.Context) {
	param := service.ArticleCreateRequest{}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid err: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 处理业务逻辑
	svc := service.NewService(ctx)
	err := svc.ArticleCreate(&param)
	if err != nil {
		global.Logger.ErrorF("svc.ArticleCreate err: %v", err)
		if err == errcode.ErrorArticleIsExistFail {
			response.ToErrorResponse(errcode.ErrorArticleIsExistFail)
		} else {
			response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		}
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新文章
// @Produce json
// @Tags 文章
// @Param id path int true "文章ID"
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Param desc body string false "文章描述" minlength(3) maxlength(100)
// @Param content body string false "文章内容" minlength(3)
// @Param cover_image_url body string false "封面图片地址" maxlength(255)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(ctx *gin.Context) {
	param := service.ArticleUpdateRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid err: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 处理逻辑
	svc := service.NewService(ctx)
	err := svc.ArticleUpdate(&param)
	if err != nil {
		global.Logger.ErrorF("svc.ArticleUpdate err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除文章
// @Produce json
// @Tags 文章
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(ctx *gin.Context) {
	param := service.ArticleDeleteRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid err: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 处理业务
	svc := service.NewService(ctx)
	err := svc.ArticleDelete(&param)
	if err != nil {
		global.Logger.ErrorF("svc.ArticleDelete err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
