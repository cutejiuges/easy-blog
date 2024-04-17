package ping

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/17 上午9:33
 * @FilePath: ping
 * @Description: 默认ping路由
 */

func Ping(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
