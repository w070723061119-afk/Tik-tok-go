package mysql

import (
	"TikTok/biz/model/user"
	"fmt"
)

func ModelMigrate() {
	if err := Db.AutoMigrate(&user.User{}); err != nil {
		fmt.Printf("user用户数据库迁移失败：%v", err)
		return
	} else {
		fmt.Printf("user用户数据库迁移成功\n")
	}

}
