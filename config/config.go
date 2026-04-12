package config

import (
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configmysql struct {
	Dsn string `yaml:"dsn"`
}

type Configredis struct {
	Addr string `yaml:"addr"`
}

func LoadConfig() (*Configmysql, *Configredis, error) {
	var configmysql Configmysql
	var configredis Configredis
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件失败: %v", err)
		return nil, nil, err
	}
	err = viper.Unmarshal(&configmysql)
	if err != nil {
		fmt.Printf("mysql配置解析失败: %v", err)
		return nil, nil, err
	}
	err = viper.Unmarshal(&configredis)
	if err != nil {
		fmt.Printf("redis配置解析失败: %v", err)
		return nil, nil, err
	}
	return &configmysql, &configredis, err
}

func initDB(config *Configmysql) (*gorm.DB, error) {
	var db *gorm.DB
	// 在这里使用config中的数据库配置来初始化数据库连接
	dsn := config.Dsn
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("初始化数据库连接失败: %v", err)
		return nil, err
	}
	return db, nil
}
func initRedis(config *Configredis) (*redis.Client, error) { // 在这里使用config中的Redis配置来初始化Redis连接
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
	return redisClient, nil
}

func LoadSQLConfig() (*gorm.DB, *redis.Client, error) {
	var configmysql *Configmysql
	var configredis *Configredis
	configmysql, configredis, err := LoadConfig()
	if err != nil {
		fmt.Printf("加载配置失败: %v", err)
		return nil, nil, err
	}

	db, err := initDB(configmysql)
	if err != nil {
		fmt.Printf("初始化数据库连接失败: %v", err)
		return nil, nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: configredis.Addr,
	})

	return db, redisClient, nil

}
