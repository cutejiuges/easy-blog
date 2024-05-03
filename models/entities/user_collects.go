package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/24 上午9:29
 * @FilePath: user_collects
 * @Description: 定义用户收藏文章实体
 */

type UserCollects struct {
	gorm.Model
	UserID    uint    `gorm:"primaryKey"`
	User      User    `gorm:"foreignKey:user_id"`
	ArticleID uint    `gorm:"primaryKey"`
	Article   Article `gorm:"foreignKey:article_id"`
}

func (u UserCollects) TableName() string {
	return "tb_user_collects"
}
