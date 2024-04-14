package global

import (
	"cutejiuges/easy-blog/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:39
 * @FilePath: global
 * @Description: 存储全局的变量
 */

var (
	Config  *conf.Config
	MysqlDB *gorm.DB
	Logger  *logrus.Logger
)
