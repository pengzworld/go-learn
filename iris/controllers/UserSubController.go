package controllers

import (
	"github.com/kataras/iris/v12"
	"go-learn/iris/config"
)

type UserSubController struct {
}

func (u *UserSubController) GetConfig(ctx iris.Context) interface{} {
	return config.C
}
