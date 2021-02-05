package service

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func NewService(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	return svc
}