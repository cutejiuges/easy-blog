package res

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 下午11:57
 * @FilePath: response
 * @Description: 封装response结构
 */

type Response struct {
	Data     any       `json:"data"`
	BaseResp *BaseResp `json:"base_resp"`
}

func NewResponse() *Response {
	return &Response{
		Data:     nil,
		BaseResp: NewBaseResp(),
	}
}

const (
	SUCCESS = 0
)

// Result 通用结果返回函数
func Result(c context.Context, ctx *app.RequestContext, resp Response) {
	ctx.JSON(consts.StatusOK, resp)
}
