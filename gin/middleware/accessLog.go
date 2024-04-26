package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

var (
	logPath = "./log"
	logFile = "gin.log"
)
var LogInstance = logrus.New()

// 日志初始化
func init() {
	// 打开文件
	logFileName := path.Join(logPath, logFile)
	// 使用滚动压缩方式记录日志
	rolling(logFileName)
	// 设置日志输出JSON格式
	//LogInstance.SetFormatter(&logrus.JSONFormatter{})
	LogInstance.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	// 设置日志记录级别
	LogInstance.SetLevel(logrus.DebugLevel)
}

func rolling(logFile string) {
	// 设置输出
	LogInstance.SetOutput(&lumberjack.Logger{
		Filename:   logFile, //日志文件位置
		MaxSize:    1,       // 单文件最大容量,单位是MB
		MaxBackups: 3,       // 最大保留过期文件个数
		MaxAge:     1,       // 保留过期文件的最大时间间隔,单位是天
		Compress:   true,    // 是否需要压缩滚动日志, 使用的 gzip 压缩
	})
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		LogInstance.WithFields(map[string]interface{}{
			"status_code": c.Writer.Status(),
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"url":         c.Request.URL.String(),
			"latency":     time.Now().Sub(startTime).String(),
			"query":       c.Request.URL.Query(),
			"ContentType": c.ContentType(),
		}).Info("access log")
	}
}
