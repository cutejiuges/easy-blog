package conf

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午2:15
 * @FilePath: logger_conf
 * @Description: 读取并存储日志的配置信息
 */

type LoggerConf struct {
	Level      string `yaml:"level"`
	Prefix     string `yaml:"prefix"`
	Directory  string `yaml:"directory"`
	ShowLine   bool   `yaml:"show_line"`
	ConsoleLog bool   `yaml:"console_log"`
	FileLog    bool   `yaml:"file_log"`
}
