package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-learn/gin/config"
	"go-learn/gin/middleware"
	"sync"
	"sync/atomic"
	"time"
)

var CloseOnce sync.Once

type Configurator func(b *gin.Engine)

func New(cfgs ...Configurator) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	b := gin.New()
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

func UseDefault(r *gin.Engine) {
	r.Use(middleware.ActiveRequest())
	r.Use(middleware.AccessLog())
	//UseLogger(r)
	UseRecovery(r)
}

func UseLogger(r *gin.Engine) {
	r.Use(gin.Logger())
}

func UseRecovery(r *gin.Engine) {
	r.Use(gin.Recovery())
}

// CloseResourcesBySignal 优雅关闭其它资源
// SIGHUP restart
// SIGTERM/SIGINT shutdown
func CloseResourcesBySignal() {
	go func() {
		for atomic.LoadInt32(&middleware.ActiveRequests) > 0 {
			time.Sleep(1000 * time.Millisecond)
		}
		CloseOnce.Do(func() {
			_ = config.DB.Close()
		})
	}()
}
