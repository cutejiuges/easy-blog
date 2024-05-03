package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 下午4:49
 * @FilePath: menu
 * @Description: 菜单表对应的实体类
 */

type Menu struct {
	gorm.Model
	MenuTitle    string   `gorm:"size:32" json:"menu_title"`                                                            //标题
	MenuTitleEN  string   `gorm:"size:32" json:"menu_title_en"`                                                         //英文标题
	Slogan       string   `gorm:"size:64" json:"slogan"`                                                                //slogan
	Abstract     Array    `gorm:"type:string" json:"abstract"`                                                          //简介
	AbstractTime int      `json:"abstract_time"`                                                                        //简介的切换时间
	Banners      []Banner `gorm:"many2many:tb_menu_banner;joinForeignKey:MenuID;JoinReferences:ImageID" json:"banners"` //图片列表
	BannerTime   int      `json:"menu_time"`                                                                            //菜单的切换时间
	Sort         int      `gorm:"size:10" json:"sort"`                                                                  //菜单的顺序
}

func (m Menu) TableName() string {
	return "tb_menu"
}
