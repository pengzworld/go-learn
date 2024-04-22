package routes

import (
	"go-learn/iris/controllers"
	"go-learn/iris/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(app *iris.Application) {
	mvc.Configure(app.Party("/user"), func(app *mvc.Application) {
		/**
		自动注入
		type UserController struct {
			UserModel *models.UserModel
		}
		*/
		app.Register(models.NewUserModel())
		app.Handle(new(controllers.UserController))
	})

}
