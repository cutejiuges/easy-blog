package conf

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:09
 * @FilePath: enter
 * @Description: 存储所有配置信息
 */

type Config struct {
	Mysql  MysqlConf  `yaml:"mysql"`
	Logger LoggerConf `yaml:"logger"`
	System SystemConf `yaml:"system"`
}
