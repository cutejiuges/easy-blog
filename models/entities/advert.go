package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 下午5:13
 * @FilePath: advert
 * @Description: 广告表的实体类
 */

type Advert struct {
	gorm.Model
	Title  string `gorm:"size:32" json:"title"` //标题
	Href   string `json:"href"`                 //跳转链接
	Images string `json:"images"`               //图片
	IsShow bool   `json:"is_show"`              //是否展示
}

func (a Advert) TableName() string {
	return "tb_advert"
}
