package config

import (
	redis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Rdb *redis.Client

func init() {
	var err error
	Db, Rdb, err = LoadSQLConfig()
	if err != nil {
		panic("加载配置失败: " + err.Error())
	}
	if err != Rdb.Ping(Rdb.Context()).Err() {
		panic("连接Redis失败: " + err.Error())
	}
}
