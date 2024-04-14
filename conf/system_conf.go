package conf

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:20
 * @FilePath: system_conf
 * @Description: 读取并存储运行时系统配置
 */

type SystemConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}
