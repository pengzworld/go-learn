package middleware

import (
	"github.com/gin-gonic/gin"
	"sync/atomic"
)

var (
	ActiveRequests int32
)

func ActiveRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt32(&ActiveRequests, 1)
		defer atomic.AddInt32(&ActiveRequests, -1)
		c.Next()
	}
}
