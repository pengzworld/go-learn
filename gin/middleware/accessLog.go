package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var logger *logrus.Logger

func init() {
	logger = newAccessLogger()
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		logger.WithFields(map[string]interface{}{
			"content_type": c.ContentType(),
			"latency":      time.Now().Sub(startTime).String(),
			"code":         c.Writer.Status(),
			"method":       c.Request.Method,
			"url":          c.Request.URL.String(),
			"client_ip":    c.ClientIP(),
		}).Info("access log")
	}
}

func newAccessLogger() *logrus.Logger {
	var logger = logrus.New()
	//
	var fWriter = &lumberjack.Logger{
		Filename:   "./log/access.log", //日志文件位置
		MaxSize:    1,                  // 单文件最大容量,单位是MB
		MaxBackups: 3,                  // 最大保留过期文件个数
		MaxAge:     1,                  // 保留过期文件的最大时间间隔,单位是天
		Compress:   false,              // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	logger.SetOutput(io.MultiWriter(fWriter, os.Stdout))
	logger.ReportCaller = true
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true, DisableHTMLEscape: true})
	return logger
}
