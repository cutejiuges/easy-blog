package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 下午5:08
 * @FilePath: fade_back
 * @Description: 用户反馈表的实体类
 */

type FadeBack struct {
	gorm.Model
	Email        string `gorm:"size:350" json:"email"`         //邮箱
	Content      string `gorm:"size:128" json:"content"`       //内容
	ApplyContent string `gorm:"size:128" json:"apply_content"` //回复内容
	IsApply      bool   `json:"is_apply"`                      //是否回复
}

func (f FadeBack) TableName() string {
	return "tb_fade_back"
}
