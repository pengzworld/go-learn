package route

import (
	"github.com/gin-gonic/gin"
	"go-learn/gin/controller"
	"go-learn/gin/middleware"
	"log"
)

func Configure(e *gin.Engine) {
	e.GET("/ping", gin.BasicAuth(gin.Accounts{"admin": "123"}), func(c *gin.Context) {
		log.Println(22222222)
		//c.Abort()
		c.Next()
	}, func(c *gin.Context) {
		log.Println(11111111)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//route group
	admin := e.Group("/admin") // 第二个参数 其实就是中间件
	{
		admin.Use(middleware.CostTime()) // 分组中间件
		admin.Use(gin.BasicAuth(gin.Accounts{"admin": "123"}))
		adminV1 := admin.Group("/v1") //路由嵌套分组
		{
			adminV1.GET("/user", func(c *gin.Context) {
				c.JSON(200, "/admin/v1/user")
			})
		}
		admin.GET("/user/:name", new(controller.UserController).Login)
	}
}
