package main

import (
	"github.com/kataras/iris/v12"
	"go-learn/iris/routes"
)

func main() {
	app := newApp()
	app.Logger().SetLevel("debug")
	_ = app.Listen(":8082", iris.WithOptimizations)
}

func newApp() *iris.Application {
	app := iris.New()
	routes.Configure(app)
	return app
}
