package middleware

import (
	"github.com/gin-gonic/gin"
	"go-learn/gin/core"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		core.AccessLogger.WithFields(map[string]interface{}{
			"content_type": c.ContentType(),
			"latency":      time.Now().Sub(startTime).String(),
			"code":         c.Writer.Status(),
			"method":       c.Request.Method,
			"url":          c.Request.URL.String(),
			"client_ip":    c.ClientIP(),
		}).Info("access log")
	}
}
