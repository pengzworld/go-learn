package controllers

import (
	"github.com/kataras/iris/v12"
	"go-learn/iris/models"
)

type UserController struct {
	UserModel *models.UserModel
}

func (u *UserController) Get(ctx iris.Context) interface{} {
	//user := models.NewUserModel()
	//ctx.Application().Logger().Info("哈哈哈哈哈")
	return u.UserModel.Find(1)
}
