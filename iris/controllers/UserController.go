package controllers

import (
	"github.com/kataras/iris/v12"
	"go-learn/iris/models"
)

var userModel = models.NewUserModel()

type UserController struct {
	UserModel *models.UserModel //自动注入
}

func (u *UserController) Get(ctx iris.Context) interface{} {
	//user := models.NewUserModel()
	//ctx.Application().Logger().Info("哈哈哈哈哈")

	return userModel.Find(1) //采用初始化方式
	//return u.UserModel.Find(1) //采用注入方式
}
