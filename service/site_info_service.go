package service

import (
	"cutejiuges/easy-blog/core"
	"cutejiuges/easy-blog/global"
	"cutejiuges/easy-blog/models/dto"
	"cutejiuges/easy-blog/models/vo"
	"cutejiuges/easy-blog/utils/converter"
	"cutejiuges/easy-blog/utils/errno"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午9:15
 * @FilePath: site_info_service
 * @Description: 站点信息的服务层操作
 */

// GetSiteInfo 查询站点信息
func GetSiteInfo() (*vo.GetSiteInfoVO, error) {
	siteInfoVO := &vo.GetSiteInfoVO{}
	converter.StructConverter(siteInfoVO, &global.Config.SiteInfo)
	global.Logger.Info("service.GetSiteInfo: ", siteInfoVO)
	return siteInfoVO, nil
}

// UpdateSiteInfo 更新站点信息
func UpdateSiteInfo(updateDTO *dto.UpdateSiteInfoDTO) (err error) {
	siteInfo := &global.Config.SiteInfo
	converter.StructConverter(siteInfo, updateDTO)
	global.Logger.Info("service.UpdateSiteInfo: ", global.Config.SiteInfo)
	bytes, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Logger.Error("service.UpdateSiteInfo, 解析yaml文件出错, err = ", err)
		return errno.NewBlogErrorWithMsg(errno.SystemError, err.Error())
	}
	err = ioutil.WriteFile(core.ConfigFile, bytes, fs.ModePerm)
	if err != nil {
		global.Logger.Error("service.UpdateSiteInfo, 写入yaml文件出错, err = ", err)
		return errno.NewBlogErrorWithMsg(errno.SystemError, err.Error())
	}
	return nil
}
