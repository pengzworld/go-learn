package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func init() {
	DB = CreateEngine()
}

func CreateEngine() *xorm.Engine {
	driver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		"remote",
		"rxLt2bdQAyf9E",
		"192.168.1.154",
		"3308",
		"hw_jiaoyu_cn",
	)
	engine, err := xorm.NewEngine("mysql", driver)
	if err != nil {
		log.Fatalf("Failed to create database engine: %v", err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(5) //最大空闲
	engine.SetMaxOpenConns(5) //最大连接
	return engine
}
