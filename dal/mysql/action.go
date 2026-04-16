package mysql

import (
	"TikTok/biz/model/user"
	"TikTok/biz/model/video"
	"fmt"

	"gorm.io/gorm"
)

func ModelMigrate() {
	if err := Db.AutoMigrate(&user.User{}); err != nil {
		fmt.Printf("user用户数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("user用户数据库迁移成功\n")
	}
	if err := Db.AutoMigrate(&video.Video{}); err != nil {
		fmt.Printf("video视频数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("video视频数据库迁移成功\n")
	}

}
func PageSelect(page, pagesize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pagesize).Limit(pagesize)
	}
}
