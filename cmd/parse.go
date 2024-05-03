package cmd

import (
	"flag"
	"os"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午1:41
 * @FilePath: parse
 * @Description: 解析命令行参数，进行分发执行
 */

type Option struct {
	DB bool
}

func Parse() Option {
	db := flag.Bool("db", false, "初始化数据库")
	flag.Parse()
	return Option{DB: *db}
}

func NeedStop(op Option) bool {
	if op.DB {
		return true
	}
	return false
}

func SwitchOption(op Option) {
	if NeedStop(op) {
		GenerateDB()
		os.Exit(0)
	}
}
