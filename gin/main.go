package main

import (
	"github.com/fvbock/endless"
	"go-learn/gin/bootstrap"
	_ "go-learn/gin/config"
	"go-learn/gin/route"
	"log"
	"syscall"

	"github.com/gin-gonic/gin"
)

// main ./web > log.log 2>&1 &
func main() {
	app := newApp()
	srv := endless.NewServer(":8083", app)

	srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGHUP] = append(
		srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGHUP],
		bootstrap.CloseResourcesBySignal)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Println("Server exiting...")
}

func newApp() *gin.Engine {
	app := bootstrap.New(bootstrap.UseDefault, route.Configure)
	return app
}
