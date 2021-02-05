package setting

import (
	"github.com/fsnotify/fsnotify"
	"time"
)

// 声明配置属性结构体

// 服务配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 应用配置
type AppSettingS struct {
	DefaultPageSize     int
	MaxPageSize         int
	LogSavePath         string
	LogFileName         string
	LogFileExt          string
	UploadSavePath      string
	UploadServerUri     string
	UploadImageMaxSize  int
	UploadImageAllowExt []string
}

// 数据库配置
type DatabaseSettingS struct {
	DBType      string
	Username    string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleTime int
	MaxOpenConn int
}

// JWT配置
type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

// 邮件配置
type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

// 中间件配置
type MiddlewareSettingS struct {
	DefaultContextTimeout time.Duration
}

var sections = make(map[string]interface{})

// 解析配置文件
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// 读取所有配置
func (s *Setting) ReloadAllSections() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// 监听配置变化
func (s *Setting) WatchSettingChange()  {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSections()
		})
	}()
}