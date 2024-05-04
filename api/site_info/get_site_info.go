package site_info

import (
	"context"
	"cutejiuges/easy-blog/service"
	"cutejiuges/easy-blog/utils"
	"cutejiuges/easy-blog/utils/res"
	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午8:51
 * @FilePath: site_info
 * @Description: 返回网站的站点信息
 */

func GetSiteInfo(c context.Context, ctx *app.RequestContext) {
	siteInfo, err := service.GetSiteInfo()
	resp := res.Response{
		Data:     siteInfo,
		BaseResp: utils.BuildBaseRespWithErr(err),
	}
	res.Result(c, ctx, resp)
}
