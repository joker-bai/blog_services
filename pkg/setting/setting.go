package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 初始化配置文件的基础属性
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp: vp}
	s.WatchSettingChange()
	return s, nil
}

