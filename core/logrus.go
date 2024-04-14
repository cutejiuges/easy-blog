package core

import (
	"bytes"
	"cutejiuges/easy-blog/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 下午3:26
 * @FilePath: logrus
 * @Description: 初始化系统日志
 */

// 日志颜色设置
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同level展示日志的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	//定义输出流
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	//定义输出格式
	//时间格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logCfg := global.Config.Logger
	if entry.HasCaller() {
		//定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//输出格式
		_, _ = fmt.Fprintf(buffer, "[%s %s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logCfg.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(buffer, "[%s %s] \x1b[%dm[%s]\x1b[0m %s\n", logCfg.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return buffer.Bytes(), nil
}

func InitLogger() {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(global.Config.Logger.ShowLine)
	mLog.SetFormatter(&LogFormatter{})
	lvl, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	mLog.SetLevel(lvl)
	InitDefaultLogger()
	global.Logger = mLog
}

func InitDefaultLogger() {
	//全局默认log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	lvl, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
}
