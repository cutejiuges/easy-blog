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
	NeedMigrateDB     bool
	NeedGenerateModel bool
}

func Parse() Option {
	migrate := flag.Bool("migrate", false, "初始化数据库")
	generate := flag.Bool("generate", false, "生成实体类")
	flag.Parse()
	return Option{
		NeedMigrateDB:     *migrate,
		NeedGenerateModel: *generate,
	}
}

func NeedStop(op Option) bool {
	if op.NeedMigrateDB || op.NeedGenerateModel {
		return true
	}
	return false
}

func SwitchOption(op Option) {
	if NeedStop(op) {
		if op.NeedMigrateDB {
			MigrateDB()
		}
		if op.NeedGenerateModel {
			GenerateModels()
		}
		os.Exit(0)
	}
}
