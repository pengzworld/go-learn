package lib

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
)

var (
	DefaultWriter io.Writer
	Logger        *logrus.Logger
)

func init() {
	newWriter()
	newLogger()
}

func newWriter() {
	DefaultWriter = &lumberjack.Logger{
		Filename:   "./log/gin.log", //日志文件位置
		MaxSize:    1,               // 单文件最大容量,单位是MB
		MaxBackups: 3,               // 最大保留过期文件个数
		MaxAge:     1,               // 保留过期文件的最大时间间隔,单位是天
		Compress:   false,           // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
}

func newLogger() {
	var logger = logrus.New()
	logger.SetOutput(DefaultWriter)
	logger.ReportCaller = true
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true, DisableHTMLEscape: true})
	Logger = logger
}
