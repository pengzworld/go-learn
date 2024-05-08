package model

import (
	. "go-learn/gin/datasource"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id              int       `xorm:"not null pk autoincr INT(11)"`
	Name            string    `xorm:"name"`
	Email           string    `xorm:"email"`
	EmailVerifiedAt time.Time `xorm:"TIMESTAMP"`
	Password        string    `xorm:"password"`
	RememberToken   string    `xorm:"remember_token"`
	CreatedAt       time.Time `xorm:"TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"TIMESTAMP"`
}

func (u *User) TableName() string {
	return "users"
}

type UserModel struct {
	Engine *xorm.Engine
}

func NewUserModel() *UserModel {
	e, _ := Database.Get("default")
	return &UserModel{
		Engine: e,
	}
}

func (u *UserModel) Find(id int) *User {
	user := &User{}
	_, _ = u.Engine.ID(id).Get(user)
	return user
}

func (u *UserModel) Last() *User {
	user := &User{}
	_, _ = u.Engine.OrderBy("id DESC").Limit(1).Get(user)
	return user
}
