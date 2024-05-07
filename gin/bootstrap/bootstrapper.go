package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-learn/gin/core"
	"go-learn/gin/datasource"
	"go-learn/gin/global/variable"
	"go-learn/gin/middleware"
	"go-learn/gin/route"
	"sync"
	"sync/atomic"
	"time"
)

var closeOnce sync.Once

func NewApplication() (app *gin.Engine) {
	core.InitConfig()
	core.InitLogger()
	//
	gin.SetMode(core.C.App.ReleaseMode)
	gin.ForceConsoleColor()
	gin.DefaultWriter = core.DefaultWriter
	gin.DefaultErrorWriter = core.DefaultWriter
	//
	app = gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.ActiveRequest(), middleware.AccessLog())
	//注册路由
	route.InitRouter(app)
	//数据库
	datasource.InitDatabase()
	return
}

// CloseResourcesBySignal 优雅关闭其它资源
// SIGHUP restart
// SIGTERM/SIGINT shutdown
func CloseResourcesBySignal() {
	go func() {
		for atomic.LoadInt32(&variable.ActiveRequests) > 0 {
			time.Sleep(1000 * time.Millisecond)
		}
		closeOnce.Do(func() {
			datasource.Database.Close()
			core.Logger.Info("closed db engines")
		})
	}()
}
