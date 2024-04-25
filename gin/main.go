package main

import (
	"log"
	"net/http"

	"go-learn/gin/bootstrap"
	"go-learn/gin/route"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func main() {
	app := newApp()
	server := []*http.Server{
		{
			Addr:    ":8083",
			Handler: app,
		},
	}
	logger := log.New(gin.DefaultWriter, "[GIN-GRACE]", 0)
	gracehttp.SetLogger(logger)
	err := gracehttp.ServeWithOptions(server, gracehttp.PreStartProcess(func() error {
		//重启前的操作,如释放外部资源等
		//kill -SIGUSR2 9347 触发重启
		logger.Println("Release other resource...")
		return nil
	}))
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func newApp() *gin.Engine {
	app := bootstrap.New(route.Configure)
	return app
}
