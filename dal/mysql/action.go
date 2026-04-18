package mysql

import (
	"TikTok/biz/model/communication"
	"TikTok/biz/model/interaction"
	"TikTok/biz/model/user"
	"TikTok/biz/model/video"
	"fmt"
	"strconv"

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
	if err := Db.AutoMigrate(&communication.Followers{}); err != nil {
		fmt.Printf("follow关注数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("follow关注数据库迁移成功\n")
	}
	if err := Db.AutoMigrate(&video.VideoLiker{}); err != nil {
		fmt.Printf("favorite点赞数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("favorite点赞数据库迁移成功\n")
	}
	if err := Db.AutoMigrate(&interaction.Comment{}); err != nil {
		fmt.Printf("comment评论数据库迁移失败：%v", err)
		return
	} else {

		fmt.Printf("comment评论数据库迁移成功\n")
	}
	if err := Db.AutoMigrate(&interaction.ChildComment{}); err != nil {
		fmt.Printf("comment评论数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("子comment评论数据库迁移成功\n")
	}

}
func PageSelect(page, pagesize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pagesize).Limit(pagesize)
	}
}

func UpdateVideoLikeCount(videoId string, likeCount string) {
	strconv.Atoi(likeCount)
	Db.Model(&video.Video{}).Where("VideoId = ?", videoId).Update("LikeCount", likeCount)
}

// UpdateVideoVisitCount 简单的数据库更新访问量函数
func UpdateVideoVisitCount(videoId string) {
	Db.Model(&video.Video{}).Where("VideoId = ?", videoId).UpdateColumn("visit_count", gorm.Expr("visit_count + ?", 1))
}
