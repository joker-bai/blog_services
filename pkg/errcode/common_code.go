package errcode

// 公共错误代码
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(100001, "服务内部错误")
	InvalidParams             = NewError(100002, "入参错误")
	NotFound                  = NewError(100003, "找不到")
	UnauthorizedAuthNotExist  = NewError(100004, "鉴权失败,找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(100005, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(100006, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(100007, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(100008, "请求太多")

	// 标签错误码
	ErrorGetTagListFail = NewError(201001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(201002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(201003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(201004, "删除标签失败")
	ErrorCountTagFail   = NewError(201005, "统计标签失败")
	ErrorTagIsExistFail = NewError(201006, "创建标签失败,标签已经存在")

	// 文件上传错误码
	ErrorUploadFileFail = NewError(203001, "上传文件失败")
)
