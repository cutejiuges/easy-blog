package main

import (
	"cutejiuges/easy-blog/cmd"
	"cutejiuges/easy-blog/core"
	"cutejiuges/easy-blog/routers"
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
	//连接数据库
	core.InitGorm()
	//解析命令行参数执行
	option := cmd.Parse()
	cmd.SwitchOption(option)
	//初始化路由
	h := routers.InitRouter()
	h.Spin()
}
