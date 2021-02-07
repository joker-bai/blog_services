package service

import (
	"code.coolops.cn/blog_services/internal/model"
)

// 定义Article请求结构体
type ArticleCountRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleCreateRequest struct {
	Title     string `form:"title" binding:"max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=10"`
}

type ArticleUpdateRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"created_by" binding:"required,min=3,max=10"`
}

type ArticleDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (s Service) ArticleList(param *ArticleListRequest) ([]*model.Article, error) {
	return s.dao.ArticleList(param.Title, param.State)
}

func (s Service) ArticleCount(param *ArticleCountRequest) (int, error) {
	return s.dao.ArticleCount(param.Title, param.State)
}

func (s Service) ArticleCreate(param *ArticleCreateRequest) error {
	return s.dao.ArticleCreate(param.Title, param.State, param.CreatedBy)
}

func (s Service) ArticleUpdate(param *ArticleUpdateRequest) error {
	return s.dao.ArticleUpdate(param.ID, param.State, param.Title, param.ModifiedBy)
}

func (s Service) ArticleDelete(param *ArticleDeleteRequest) error {
	return s.dao.ArticleDelete(param.ID)
}
