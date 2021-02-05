package global

// 全局配置文件
import "code.coolops.cn/blog_services/pkg/setting"

var (
	ServerSetting     *setting.ServerSettingS
	AppSetting        *setting.AppSettingS
	DatabaseSetting   *setting.DatabaseSettingS
	JWTSetting        *setting.JWTSettingS
	EmailSetting      *setting.EmailSettingS
	MiddlewareSetting *setting.MiddlewareSettingS
)
