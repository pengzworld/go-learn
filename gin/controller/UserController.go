package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-learn/gin/config"
	"os"
	"time"
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

type UserController struct {
}

func (u *UserController) Login(c *gin.Context) {
	engine := config.DB

	duration, _ := time.ParseDuration(c.DefaultQuery("time", "0s"))
	time.Sleep(duration)

	user := &User{}
	_, _ = engine.ID(1).Get(user)

	//a := []int{1, 0}
	//for _, v := range a {
	//	b := 10 / v
	//	fmt.Println(b)
	//}

	c.JSON(200, gin.H{
		"user": user,
		"pid":  os.Getpid(),
	})

	//a := []int{1, 0}
	//for _, v := range a {
	//	b := 10 / v
	//	fmt.Println(b)
	//}
	//name := c.Param("name")
	//q := c.Query("q")
	//m := c.DefaultQuery("q", "0")
	//n, _ := c.GetQueryArray("q[]")
	//c.QueryArray  接受数组 ?media=blog&media=wechat
	//c.QueryMap("ids") 接受map    ?ids[a]=123&ids[b]=456&ids[c]=789
	//c.JSON(200, gin.H{
	//	"a": name,
	//	"b": q,
	//	"c": m,
	//	"n": n,
	//})
}
