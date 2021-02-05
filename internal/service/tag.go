package service

import (
	"code.coolops.cn/blog_services/internal/model"
	"code.coolops.cn/blog_services/pkg/app"
)

// 定义Tag请求结构体
type TagCountRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagCreateRequest struct {
	Name      string `form:"name" binding:"max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=10"`
}

type TagUpdateRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=10"`
}

type TagDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (s Service) TagCount(param *TagCountRequest) (int, error) {
	return s.dao.TagCount(param.Name, param.State)
}

func (s Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.TagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s Service) TagCreate(param *TagCreateRequest) error {
	return s.dao.TagCreate(param.Name, param.State, param.CreatedBy)
}

func (s Service) TagUpdate(param *TagUpdateRequest) error {
	return s.dao.TagUpdate(param.ID, param.State, param.Name, param.ModifiedBy)
}

func (s Service) TagDelete(param *TagDeleteRequest) error {
	return s.dao.TagDelete(param.ID)
}
