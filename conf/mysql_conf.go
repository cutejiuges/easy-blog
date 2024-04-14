package conf

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:11
 * @FilePath: mysql_conf
 * @Description: 读取并存储MySQL的配置信息
 */

type MysqlConf struct {
	Host         string `yaml:"host"`           //服务器地址
	Port         int    `yaml:"port"`           //服务端口
	DB           string `yaml:"db"`             //db name
	UserName     string `yaml:"user_name"`      //用户名
	Password     string `yaml:"password"`       //密码
	Config       string `yaml:"config"`         //mysql高级设置，如字符集、排序规则等
	MaxIdleConns int    `yaml:"max_idle_conns"` //最大空闲连接数
	MaxOpenConns int    `yaml:"max_open_conns"` //最大容纳连接数
	MaxLifeTime  int    `yaml:"max_life_time"`  //连接最大复用时间，单位为秒
	LogLevel     string `yaml:"log_level"`      //日志等级
}

func (m *MysqlConf) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
