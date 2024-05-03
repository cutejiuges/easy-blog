package cmd

import (
	"cutejiuges/easy-blog/global"
	"cutejiuges/easy-blog/models/entities"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午1:49
 * @FilePath: generate_db
 * @Description: 根据实体类模型生成数据库表结构
 */

func GenerateDB() {
	var err error
	global.MysqlDB.SetupJoinTable(&entities.User{}, "CollectArticles", &entities.UserCollects{})
	global.MysqlDB.SetupJoinTable(&entities.Menu{}, "Banners", &entities.MenuBanner{})
	err = global.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&entities.User{},
			&entities.Banner{},
			&entities.Tag{},
			&entities.Message{},
			&entities.Advert{},
			&entities.Comment{},
			&entities.Article{},
			&entities.Menu{},
			&entities.MenuBanner{},
			&entities.FadeBack{},
			&entities.LoginData{},
		)
	if err != nil {
		global.Logger.Errorf("[error] 生成表结构失败了, err = %v", err)
		return
	}
	global.Logger.Info("[success] 生成表结构成功!")
}
