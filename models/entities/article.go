package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/24 上午12:30
 * @FilePath: article
 * @Description: 存储文章的实体类结构
 */

type Article struct {
	gorm.Model
	Title        string    `gorm:"size:32" json:"title"`                 //文章标题
	Abstract     string    `json:"abstract"`                             //文章简介
	Content      string    `json:"content"`                              //文章内容
	LookCount    int       `json:"look_count"`                           //浏览量
	CommentCount int       `json:"comment_count"`                        //评论量
	DiggCount    int       `json:"digg_count"`                           //获赞量
	CollectCount int       `json:"collect_count"`                        //收藏量
	Tags         []Tag     `gorm:"many2many:tb_article_tag" json:"tags"` //文章标签
	Comments     []Comment `gorm:"foreignKey:article_id" json:"-"`       //文章评论
	User         User      `gorm:"foreignKey:user_id" json:"-"`          //用户信息
	UserID       uint      `json:"user_id"`                              //用户id
	Category     string    `gorm:"size:20" json:"category"`              //文章分类
	Source       string    `json:"source"`                               //文章来源
	Link         string    `json:"link"`                                 //原文链接
	Banner       Banner    `gorm:"foreignKey:banner_id" json:"-"`        //文章封面
	BannerID     uint      `json:"banner_id"`                            //封面id
	NickName     string    `gorm:"size:36" json:"nick_name"`             //作者昵称
	BannerPath   string    `json:"banner_path"`                          //文章封面
	TagArray     Array     `gorm:"type:string;size:64" json:"tag_array"` //文章标签
}

func (Article) TableName() string {
	return "tb_article"
}
