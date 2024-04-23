package main

import (
	"go-learn/iris/bootstrap"
	"go-learn/iris/middleware/access_log"
	"go-learn/iris/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()
	_ = app.Listen(":8082", iris.WithOptimizations)
}

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New(bootstrap.SetLogger)
	app.Bootstrap()
	app.Configure(access_log.Configure, routes.Configure)
	return app
}
