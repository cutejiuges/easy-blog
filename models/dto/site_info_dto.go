package dto

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午6:19
 * @FilePath: site_info_dto
 * @Description: 封装site_info相关的dto
 */

// UpdateSiteInfoDTO 更新站点信息的dto
type UpdateSiteInfoDTO struct {
	CreatedAt  string `json:"created_at"`
	FilingInfo string `json:"filing_info"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Job        string `json:"job"`
	Addr       string `json:"addr"`
	Slogan     string `json:"slogan"`
	SloganEN   string `json:"slogan_en"`
	GithubUrl  string `json:"github_url"`
}
