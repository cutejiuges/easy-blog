package entities

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 下午5:00
 * @FilePath: menu_image
 * @Description: 菜单和背景的连接表
 */

type MenuBanner struct {
	MenuID   uint   `json:"menu_id"`             //关联的菜单id
	Menu     Menu   `gorm:"foreignKey:MenuID"`   //菜单
	BannerID uint   `json:"banner_id"`           //关联的图片id
	Banner   Banner `gorm:"foreignKey:BannerID"` //图片
	Sort     int    `gorm:"size:10" json:"sort"` //排序
}

func (m MenuBanner) TableName() string {
	return "tb_menu_banner"
}
