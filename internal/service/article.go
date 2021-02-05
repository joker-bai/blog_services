package service

// 定义Article请求结构体
type ArticleCountRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8	`form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8	`form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleCreateRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8	`form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=10"`
}

type ArticleUpdateRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Title string `form:"title" binding:"max=100"`
	State uint8	`form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=10"`
}

type ArticleDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
