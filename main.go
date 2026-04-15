package main

import (
	"TikTok/dal/mysql"
	myredis "TikTok/dal/redis"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// 初始化数据库和 Redis

	h := server.Default(server.WithHostPorts("0.0.0.0:8080"))
	if err := mysql.Init(); err != nil {
		panic("Failed to initialize MySQL: " + err.Error())
	}
	if err := myredis.Init(); err != nil {
		panic("Failed to initialize Redis: " + err.Error())
	}
	register(h)
	h.Spin()
}
