package middleware

import (
	"sync/atomic"

	"go-learn/gin/global/variable"

	"github.com/gin-gonic/gin"
)

func ActiveRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt32(&variable.ActiveRequests, 1)
		defer atomic.AddInt32(&variable.ActiveRequests, -1)
		//bootstrap.Logger.Info("ActiveRequests: 前")
		c.Next()
		//bootstrap.Logger.Info("ActiveRequests: 后")
	}
}
