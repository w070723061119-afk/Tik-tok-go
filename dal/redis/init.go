package myredis

import (
	"fmt"

	"TikTok/config"

	redis "github.com/go-redis/redis/v8"
)

var (
	Rdb          *redis.Client
	configLoaded bool
)

func initRedis(cfg *config.Configredis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	return redisClient, nil
}

func Init() error {
	if configLoaded {
		return nil
	}

	redisConfig, err := config.LoadRedisConfig()
	if err != nil {
		fmt.Printf("加载 Redis 配置失败：%v", err)
		return err
	}

	Rdb, err = initRedis(redisConfig)
	if err != nil {
		fmt.Printf("初始化 Redis 连接失败：%v", err)
		return err
	}

	err = Rdb.Ping(Rdb.Context()).Err()
	if err != nil {
		fmt.Printf("连接 Redis 失败：%v", err)
		return err
	}

	configLoaded = true
	return nil
}

func GetRedis() *redis.Client {
	return Rdb
}
