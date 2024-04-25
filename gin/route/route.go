package route

import (
	"github.com/gin-gonic/gin"
	"go-learn/gin/controller"
	"go-learn/gin/middleware"
)

func Configure(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//route group
	admin := e.Group("/admin") // 第二个参数 其实就是中间件
	{
		admin.Use(middleware.CostTime()) // 分组中间件
		adminV1 := admin.Group("/v1")    //路由嵌套分组
		{
			adminV1.GET("/user", func(c *gin.Context) {
				c.JSON(200, "/admin/v1/user")
			})
		}
		admin.GET("/user/:name", new(controller.UserController).Login)
	}
}
