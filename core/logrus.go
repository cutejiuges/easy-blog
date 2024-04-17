package core

import (
	"bytes"
	"cutejiuges/easy-blog/global"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
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
		//运行时文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//输出格式
		_, _ = fmt.Fprintf(buffer, "[%s] [%s] [%s] %s %s %s\n", logCfg.Prefix, timestamp, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(buffer, "[%s] [%s] [%s] %s\n", logCfg.Prefix, timestamp, entry.Level, entry.Message)
	}
	return buffer.Bytes(), nil
}

func logFileCut(fileName string) *rotatelogs.RotateLogs {
	logier, err := rotatelogs.New(
		//切割后日志文件名
		fileName,
		//日志文件最大的保存时间
		rotatelogs.WithMaxAge(30*24*time.Hour),
		//日志切割的时间间隔
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	logFileHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logier,
		logrus.FatalLevel: logier,
		logrus.DebugLevel: logier,
		logrus.WarnLevel:  logier,
		logrus.ErrorLevel: logier,
		logrus.PanicLevel: logier,
	}, &LogFormatter{})
	logrus.AddHook(logFileHook)
	return logier
}

func InitLogger() {
	//按时间进行切割
	err := os.MkdirAll("/home/wujie/myFiles/GoCode/easy-blog/log", 0755)
	if err != nil {
		panic(err)
	}
	logFileName := path.Join("/home/wujie/myFiles/GoCode/easy-blog/log", "easy_blog") + ".%Y%m%d.log"
	logFile := logFileCut(logFileName)

	mLog := logrus.New()
	//以自定义格式打日志
	mLog.SetFormatter(&LogFormatter{})

	//是否输出到命令行和日志文件
	writers := []io.Writer{}
	if global.Config.Logger.ConsoleLog {
		writers = append(writers, os.Stdout)
	}
	if global.Config.Logger.FileLog {
		writers = append(writers, logFile)
	}
	mLog.SetOutput(io.MultiWriter(writers...))

	//显示文件名、函数名和行号
	mLog.SetReportCaller(global.Config.Logger.ShowLine)

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
