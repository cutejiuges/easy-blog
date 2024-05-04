package core

import (
	"cutejiuges/easy-blog/global"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 下午2:06
 * @FilePath: gorm
 * @Description: 此处完成gorm的连接和获取
 */

var once sync.Once

type myWriter struct {
	mLog *logrus.Logger
}

// Printf 实现gorm/logger.Writer接口
func (m *myWriter) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	//记录日志
	m.mLog.Info(logStr)
}

func newMyWriter() *myWriter {
	return &myWriter{mLog: global.Logger}
}

func InitGorm() {
	var (
		db  *gorm.DB
		err error
	)
	once.Do(func() {
		db, err = mysqlConnect()
	})
	if err != nil {
		panic(err)
	}
	global.MysqlDB = db
	global.Logger.Info("db初始化连接成功,", global.MysqlDB)
}

func mysqlConnect() (*gorm.DB, error) {
	if global.Config.Mysql.Host == "" {
		global.Logger.Warn("未配置Mysql，取消mysql连接")
		return nil, errors.New("mysql未配置")
	}

	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		//开发环境打印所有的sql语句
		mysqlLogger = logger.New(newMyWriter(), logger.Config{
			LogLevel: logger.Info,
		})
	} else {
		//其他环境只打印警告及以上的sql语句
		mysqlLogger = logger.New(newMyWriter(), logger.Config{
			LogLevel: logger.Warn,
		})
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logger.Fatal(fmt.Sprintf("[%s] mysql连接失败!", dsn))
	}
	MysqlDB, _ := db.DB()
	MysqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	MysqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	MysqlDB.SetConnMaxLifetime(time.Second * time.Duration(global.Config.Mysql.MaxIdleConns))
	return db, nil
}
