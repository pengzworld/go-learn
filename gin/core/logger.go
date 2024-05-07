package core

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
)

var (
	DefaultWriter io.Writer
	Logger        *logrus.Logger
)

func InitLogger() {
	DefaultWriter = &lumberjack.Logger{
		Filename:   "./log/gin.log", //日志文件位置
		MaxSize:    1,               // 单文件最大容量,单位是MB
		MaxBackups: 3,               // 最大保留过期文件个数
		MaxAge:     1,               // 保留过期文件的最大时间间隔,单位是天
		Compress:   false,           // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	Logger = logrus.New()
	Logger.SetOutput(DefaultWriter)
	Logger.ReportCaller = true
	Logger.SetLevel(C.App.LogLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true, DisableHTMLEscape: true})
}
