package v1

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/service"
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/convert"
	"code.coolops.cn/blog_services/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Tags 标签
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(ctx *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	// 业务逻辑处理
	svc := service.NewService(ctx)
	pager := app.Pager{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPageSize(ctx),
	}
	totalRows, err := svc.TagCount(&service.TagCountRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		global.Logger.ErrorF("svc.TagCount err:%v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.ErrorF("svc.GetTagList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增标签
// @Produce json
// @Tags 标签
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(ctx *gin.Context) {
	param := service.TagCreateRequest{}
	response := app.NewResponse(ctx)
	valied, errors := app.BindAndValid(ctx, &param)
	if !valied {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 进行插入操作
	svc := service.NewService(ctx)
	err := svc.TagCreate(&param)
	if err != nil {
		global.Logger.ErrorF("svc.TagCreate err: %v", err)
		if err == errcode.ErrorTagIsExistFail {
			response.ToErrorResponse(errcode.ErrorTagIsExistFail)
		} else {
			response.ToErrorResponse(errcode.ErrorCreateTagFail)
		}
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Tags 标签
// @Param id path int true "标签ID"
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(ctx *gin.Context) {
	param := service.TagUpdateRequest{
		ID: convert.StrTo(ctx.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 更新操作
	svc := service.NewService(ctx)
	err := svc.TagUpdate(&param)
	if err != nil {
		global.Logger.ErrorF("svc.TagUpdate errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Tags 标签
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(ctx *gin.Context) {
	param := service.TagDeleteRequest{
		ID: convert.StrTo(ctx.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errors)
		errRsp := errcode.InvalidParams.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	// 删除操作
	svc := service.NewService(ctx)
	err := svc.TagDelete(&param)
	if err != nil {
		global.Logger.ErrorF("svc.TagDelete errs: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
