package main

import (
	"cutejiuges/easy-blog/core"
	"cutejiuges/easy-blog/global"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/13 下午6:26
 * @FilePath: main
 * @Description: 程序入口文件
 */
func main() {
	//读取配置文件
	core.InitConf()
	//初始化系统日志
	core.InitLogger()
	global.Logger.Info("111")
	global.Logger.Warn("222")
	global.Logger.Error("333")
	//连接数据库
	core.InitGorm()
}
