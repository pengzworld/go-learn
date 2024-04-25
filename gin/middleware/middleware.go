package middleware

import "github.com/gin-gonic/gin"

func Configure(e *gin.Engine) {
	e.Use(gin.Logger(), gin.Recovery())
}
