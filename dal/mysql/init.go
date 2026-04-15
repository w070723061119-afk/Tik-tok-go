package mysql

import (
	"fmt"

	"TikTok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db           *gorm.DB
	configLoaded bool
)

func initDB(cfg *config.Configmysql) (*gorm.DB, error) {
	var db *gorm.DB
	dsn := cfg.Dsn
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("初始化数据库连接失败：%v", err)
		return nil, err
	}
	return db, nil
}

func Init() error {
	if configLoaded {
		return nil
	}

	mysqlConfig, err := config.LoadMysqlConfig()
	if err != nil {
		fmt.Printf("加载 MySQL 配置失败：%v", err)
		return err
	}

	Db, err = initDB(mysqlConfig)
	if err != nil {
		fmt.Printf("初始化数据库连接失败：%v", err)
		return err
	}
	ModelMigrate()
	configLoaded = true
	return nil
}

func GetDB() *gorm.DB {
	return Db
}
