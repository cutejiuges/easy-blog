package vo

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午7:03
 * @FilePath: site_info_vo
 * @Description: site_info的视图结构
 */

type GetSiteInfoVO struct {
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
