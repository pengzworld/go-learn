package routes

import (
	"go-learn/iris/bootstrap"
	"go-learn/iris/controllers"
	"go-learn/iris/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(app *bootstrap.Bootstrapper) {
	app.PartyFunc("/path1", func(r iris.Party) {
		r.Use(func(ctx iris.Context) {
			ctx.Application().Logger().Info("path1 中间件")
			ctx.Next()
		})
		r.PartyFunc("/path2", func(r iris.Party) {
			r.Use(func(ctx iris.Context) {
				ctx.Application().Logger().Info("path2 中间件")
				ctx.Next()
			})
			mvc.Configure(r, func(m *mvc.Application) {
				/**
				自动注入
				type UserController struct {
					UserModel *models.UserModel
				}
				*/
				m.Register(models.NewUserModel())
				m.Handle(new(controllers.UserSubController)) //http://localhost:8082/path1/path2/config
			})
		})

		mvc.Configure(r, func(m *mvc.Application) {
			/**
			自动注入
			type UserController struct {
				UserModel *models.UserModel
			}
			*/
			m.Register(models.NewUserModel())
			m.Handle(new(controllers.UserController)) //http://localhost:8082/path1
		})
	})
}
