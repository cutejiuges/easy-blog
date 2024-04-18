package utils

import (
	"cutejiuges/easy-blog/utils/errno"
	"cutejiuges/easy-blog/utils/res"
	"errors"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/18 上午9:15
 * @FilePath: build_base_resp
 * @Description: 通过error信息构造baseResp
 */

func BuildBaseRespWithErr(err error) *res.BaseResp {
	if err == nil {
		return &res.BaseResp{
			StatusCode:    int32(errno.SUCCESS),
			StatusMessage: errno.ErrDesc[errno.SUCCESS],
		}
	}
	var t *errno.BlogError
	switch {
	case errors.As(err, &t):
		resp := &res.BaseResp{
			StatusCode:    int32(t.Code),
			StatusMessage: t.Msg,
		}
		if resp.StatusMessage == "" {
			resp.StatusMessage = errno.ErrDesc[resp.StatusCode]
		}
		return resp
	default:
		return &res.BaseResp{
			StatusCode:    errno.SystemError,
			StatusMessage: errno.ErrDesc[errno.SystemError],
		}
	}
}
