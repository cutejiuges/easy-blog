package site_info

import (
	"context"
	"cutejiuges/easy-blog/models/dto"
	"cutejiuges/easy-blog/service"
	"cutejiuges/easy-blog/utils"
	"cutejiuges/easy-blog/utils/errno"
	"cutejiuges/easy-blog/utils/res"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午3:08
 * @FilePath: update_site_info
 * @Description: 更新站点信息
 */

func UpdateSiteInfo(c context.Context, ctx *app.RequestContext) {
	//获取请求中的更新参数
	newSiteInfo := &dto.UpdateSiteInfoDTO{}
	err := ctx.BindJSON(newSiteInfo)
	if err != nil {
		res.Result(c, ctx, res.Response{
			BaseResp: utils.BuildBaseRespWithErr(errno.NewBlogErrorWithMsg(errno.ParamIllegal, err.Error())),
		})
		return
	}
	//参数校验
	err = checkParam(newSiteInfo)
	if err != nil {
		res.Result(c, ctx, res.Response{
			BaseResp: utils.BuildBaseRespWithErr(err),
		})
		return
	}
	//更新站点信息
	err = service.UpdateSiteInfo(newSiteInfo)
	res.Result(c, ctx, res.Response{
		BaseResp: utils.BuildBaseRespWithErr(err),
	})
}

func checkParam(siteInfo *dto.UpdateSiteInfoDTO) (err error) {
	if siteInfo == nil {
		err = errno.NewBlogErrorWithMsg(errno.ParamIllegal, "请求参数解析为空")
		return err
	}
	if len(strings.TrimSpace(siteInfo.GithubUrl)) == 0 {
		err = errno.NewBlogErrorWithMsg(errno.ParamIllegal, "github链接不允许为空")
		return err
	}
	if len(strings.TrimSpace(siteInfo.Title)) == 0 {
		err = errno.NewBlogErrorWithMsg(errno.ParamIllegal, "站点标题不允许为空")
		return err
	}
	if len(strings.TrimSpace(siteInfo.Email)) == 0 {
		err = errno.NewBlogErrorWithMsg(errno.ParamIllegal, "网页邮箱不允许为空")
		return err
	}
	return err
}
