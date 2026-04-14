package token

import (
	"TikTok/config"
)

var Jwtconfig *config.Configjwt

func init() {
	var err error
	Jwtconfig, err = config.LoadJwtConfig()
	if err != nil {
		panic("加载 JWT 配置失败：" + err.Error())
	}
}
