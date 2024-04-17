package api

import (
	"cutejiuges/easy-blog/api/ping"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/17 上午9:31
 * @FilePath: router_register
 * @Description: 全局路由注册器
 */

func RegisterRouter(router *server.Hertz) {
	router.GET("/ping", ping.Ping)
}
