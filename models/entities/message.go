package entities

import "gorm.io/gorm"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 下午5:16
 * @FilePath: message
 * @Description: 消息表的实体类
 */

type Message struct {
	gorm.Model
	SendUserID          uint   `gorm:"primaryKey" json:"send_user_id"` //发送人相关信息
	SendUser            User   `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName    string `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar      string `json:"send_user_avatar"`
	ReceiveUserID       uint   `gorm:"primaryKey" json:"receive_user_id"` //接收人相关信息
	ReceiveUser         User   `gorm:"foreignKey:ReceiveUserID" json:"-"`
	ReceiveUserNickName string `gorm:"size:42" json:"receive_user_nick_name"`
	ReceiveUserAvatar   string `json:"receive_user_avatar"`
	IsRead              bool   `gorm:"default:false" json:"is_read"` //是否已读
	Content             string `json:"content"`                      //消息内容
}

func (m Message) TableName() string {
	return "tb_message"
}
