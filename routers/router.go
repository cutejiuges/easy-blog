package routers

import (
	"cutejiuges/easy-blog/api"
	"cutejiuges/easy-blog/global"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/17 上午9:18
 * @FilePath: router
 * @Description: 初始化全局路由
 */

func InitRouter() *server.Hertz {
	h := server.Default(server.WithHostPorts(global.Config.System.Addr()))
	api.RegisterRouter(h)
	return h
}
