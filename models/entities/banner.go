package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/24 上午8:57
 * @FilePath: image
 * @Description: 存放图片的实体类结构
 */

type Banner struct {
	gorm.Model
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片的hash值，判重
	Name string `gorm:"size:38" json:"name"` //图片名称
}

func (Banner) TableName() string {
	return "tb_banner"
}
