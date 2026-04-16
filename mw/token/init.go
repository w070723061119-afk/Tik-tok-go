package token

import (
	"TikTok/config"
	"os"

	"github.com/spf13/viper"
)

var Jwtconfig *config.Configjwt

func init() {
	// 优先查找根目录的 config.yml
	if _, err := os.Stat("config.yml"); err == nil {
		viper.SetConfigFile("config.yml")
	} else {
		viper.SetConfigFile("config/config.yml")
	}
	viper.SetConfigType("yaml")

	var err error
	Jwtconfig, err = config.LoadJwtConfig()
	if err != nil {
		panic("加载 JWT 配置失败：" + err.Error())
	}
	
	// 初始化 JWT 中间件
	NewAuthMiddleware()
}
