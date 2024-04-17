package conf

import "fmt"

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

func (s SystemConf) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
