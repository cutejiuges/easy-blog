package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/24 上午9:11
 * @FilePath: tag
 * @Description: 文章标签的实体类
 */

type Tag struct {
	gorm.Model
	Title    string    `gorm:"size:16" json:"title"`              //标签的名称
	Articles []Article `gorm:"many2many:tb_article_tag" json:"-"` //关联该标签的文章列表
}

func (Tag) TableName() string {
	return "tb_tag"
}
