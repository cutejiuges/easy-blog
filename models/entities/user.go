package entities

import (
	"gorm.io/gorm"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/23 下午11:39
 * @FilePath: user
 * @Description: 存储用户的实体类结构
 */

type User struct {
	gorm.Model
	NickName        string         `gorm:"size:36" json:"nick_name"`                                                        //昵称
	UserName        string         `gorm:"size:36" json:"user_name"`                                                        //用户名
	Password        string         `gorm:"size:128" json:"password"`                                                        //用户密码
	Avatar          string         `gorm:"size:255" json:"avatar"`                                                          //用户头像
	Email           string         `gorm:"size:350" json:"email"`                                                           //用户邮箱
	Tel             string         `gorm:"size:18" json:"tel"`                                                              //用户电话
	Address         string         `gorm:"size:128" json:"address"`                                                         //用户地址
	Token           string         `gorm:"size:64" json:"token"`                                                            //其他平台的唯一id
	IP              string         `gorm:"size:20" json:"ip"`                                                               //ip地址
	Role            RoleType       `gorm:"size:4;default:2" json:"role"`                                                    //用户角色
	SignOrigin      SignOriginType `gorm:"type=smallint(6)" json:"sign_origin"`                                             //注册来源
	Articles        []Article      `gorm:"foreignKey:UserID" json:"-"`                                                      //发布的文章列表
	CollectArticles []Article      `gorm:"many2many:tb_user_collects;joinForeign:UserID;joinReferences:ArticleID" json:"-"` //收藏的文章列表
}

func (User) TableName() string {
	return "tb_user"
}
