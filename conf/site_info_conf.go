package conf

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午8:42
 * @FilePath: site_info_conf
 * @Description: 站点信息的配置
 */

type SiteInfoConf struct {
	CreatedAt  string `yaml:"created_at"`
	FilingInfo string `yaml:"filing_info"`
	Title      string `yaml:"title"`
	Email      string `yaml:"email"`
	Name       string `yaml:"name"`
	Job        string `yaml:"job"`
	Addr       string `yaml:"addr"`
	Slogan     string `yaml:"slogan"`
	SloganEN   string `yaml:"slogan_en"`
	GithubUrl  string `yaml:"github_url"`
}
