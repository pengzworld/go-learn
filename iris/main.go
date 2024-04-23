package main

import (
	"fmt"

	"go-learn/iris/bootstrap"
	"go-learn/iris/config"
	"go-learn/iris/middleware/access_log"
	"go-learn/iris/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()
	addr := fmt.Sprintf("%s:%d", config.C.App.IP, config.C.App.Port)
	_ = app.Listen(addr, iris.WithConfiguration(config.C.Iris))
}

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New(bootstrap.SetLogger)
	app.Bootstrap()
	app.Configure(access_log.Configure, routes.Configure)
	return app
}
