package core

import (
	"cutejiuges/easy-blog/conf"
	"cutejiuges/easy-blog/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:27
 * @FilePath: conf
 * @Description: 读取yaml文件的配置
 */

const (
	ConfigFile = "settings.yaml"
)

func InitConf() {
	cfg := &conf.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错, err = %v", err))
	}
	err = yaml.Unmarshal(yamlConf, cfg)
	if err != nil {
		panic(fmt.Errorf("解析配置文件出错, err = %v", err))
	}
	log.Println("载入配置文件成功!")
	global.Config = cfg
	fmt.Println(global.Config)
}
