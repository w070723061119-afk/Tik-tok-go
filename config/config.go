package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configmysql struct {
	Dsn string `yaml:"dsn"`
}

type Configredis struct {
	Addr string `yaml:"addr"`
}

var configLoaded bool

type Configjwt struct {
	Jwtsecret          string `yaml:"jwtsecret"`
	Accesssecret       string `yaml:"accesssecret"`
	Refreshtokenexpire int    `yaml:"refreshtokenexpire"`
	Accesstokenexpire  int    `yaml:"accesstokenexpire"`
	Identitykey        string `yaml:"identitykey"`
}

func LoadJwtConfig() (*Configjwt, error) {
	var configjwt Configjwt
	if !configLoaded {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("读取配置文件失败: %v", err)
			return nil, err
		}
		configLoaded = true
	}
	err := viper.UnmarshalKey("jwt", &configjwt)
	if err != nil {
		fmt.Printf("jwt配置解析失败: %v", err)
		return nil, err
	}
	return &configjwt, nil
}

func LoadMysqlConfig() (*Configmysql, error) {
	if !configLoaded {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("读取配置文件失败：%v", err)
			return nil, err
		}
		configLoaded = true
	}

	var configmysql Configmysql
	err := viper.UnmarshalKey("mysql", &configmysql)
	if err != nil {
		fmt.Printf("mysql 配置解析失败：%v", err)
		return nil, err
	}
	return &configmysql, nil
}

func LoadRedisConfig() (*Configredis, error) {

	if !configLoaded {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("读取配置文件失败：%v", err)
			return nil, err
		}
		configLoaded = true
	}

	var configredis Configredis
	err := viper.UnmarshalKey("redis", &configredis)
	if err != nil {
		fmt.Printf("redis 配置解析失败：%v", err)
		return nil, err
	}
	return &configredis, nil
}
