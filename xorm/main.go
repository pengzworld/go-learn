package main

import (
	"go-learn/xorm/models"
	"time"
	"xorm.io/builder"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var users []*models.User
	userModel := models.NewUserModel()
	fields := []string{"id", "name", "created_at"}
	startTime, err := time.Parse("2006-01-02 15:04:05", "2024-04-18 09:22:47")
	if err != nil {
		panic(err)
	}
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"name": "Test User"}.Or(builder.Eq{"email": "test@example.com"}))
	cond = cond.And(builder.Gte{"created_at": startTime})
	//cond = cond.And(builder.Lte{"created_at": time.Now()})
	err = userModel.Engine.Cols(fields...).Where(cond).Limit(10, 0).OrderBy("id DESC").Find(&users)
	if err != nil {
		panic(err)
	}
	spew.Dump(users)
}
