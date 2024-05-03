package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午2:06
 * @FilePath: login_data
 * @Description: 用户登录数据
 */

type LoginData struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
	IP       string `gorm:"size:20" json:"ip"`
	NickName string `gorm:"size:42" json:"nick_name"`
	Token    string `gorm:"size:256" json:"token"`
	Device   string `gorm:"size:256" json:"device"`
	Addr     string `gorm:"size:64" json:"addr"`
}

func (l LoginData) TableName() string {
	return "tb_login_data"
}
