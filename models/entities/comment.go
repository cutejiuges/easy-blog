package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/24 上午9:15
 * @FilePath: comment
 * @Description: 用户评论的实体类结构
 */

type Comment struct {
	gorm.Model
	SubComments     []*Comment `gorm:"foreignKey:parent_comment_id" json:"sub_comments"`    //子评论列表
	ParentComments  *Comment   `gorm:"foreignKey:parent_comment_id" json:"parent_comments"` //父级评论
	ParentCommentID *uint      `json:"parent_comment_id"`                                   //父级评论id
	Content         string     `gorm:"size:256" json:"content"`                             //评论内容
	DiggCount       int        `gorm:"size:8;default:0;" json:"digg_count"`                 //获赞量
	CommentCount    int        `gorm:"size:8;default:0;" json:"comment_count"`              //评论回复量
	Article         Article    `gorm:"foreignKey:article_id" json:"article"`                //关联的文章
	ArticleID       uint       `json:"article_id"`                                          //关联文章id
	User            User       `json:"user"`                                                //关联的用户
	UserID          uint       `json:"user_id"`                                             //关联用户id
}

func (c Comment) TableName() string {
	return "tb_comment"
}
