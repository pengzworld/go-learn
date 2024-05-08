package core

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	DefaultWriter io.Writer
	Logger        *logrus.Logger
	AccessLogger  *logrus.Logger
)

func InitLogger() {
	newLogger()
	newAccessLogger()
}

func newLogger() {
	DefaultWriter = &lumberjack.Logger{
		Filename:   C.App.Log.Path,       //日志文件位置
		MaxSize:    C.App.Log.MaxSize,    // 单文件最大容量,单位是MB
		MaxBackups: C.App.Log.MaxBackups, // 最大保留过期文件个数
		MaxAge:     C.App.Log.MaxAge,     // 保留过期文件的最大时间间隔,单位是天
		Compress:   C.App.Log.Compress,   // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	Logger = logrus.New()
	Logger.SetOutput(DefaultWriter)
	Logger.ReportCaller = true
	Logger.SetLevel(C.App.Log.Level)
	Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: C.App.Log.PrettyPrint,
		DisableHTMLEscape: C.App.Log.DisableHTMLEscape})
}

func newAccessLogger() {
	AccessLogger = logrus.New()
	//
	var fWriter = &lumberjack.Logger{
		Filename:   C.App.AccessLog.Path,       //日志文件位置
		MaxSize:    C.App.AccessLog.MaxSize,    // 单文件最大容量,单位是MB
		MaxBackups: C.App.AccessLog.MaxBackups, // 最大保留过期文件个数
		MaxAge:     C.App.AccessLog.MaxAge,     // 保留过期文件的最大时间间隔,单位是天
		Compress:   C.App.AccessLog.Compress,   // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	AccessLogger.SetOutput(io.MultiWriter(fWriter, os.Stdout))
	AccessLogger.ReportCaller = C.App.AccessLog.ReportCaller
	AccessLogger.SetLevel(C.App.AccessLog.Level)
	AccessLogger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: C.App.AccessLog.PrettyPrint,
		DisableHTMLEscape: C.App.AccessLog.DisableHTMLEscape})
}
