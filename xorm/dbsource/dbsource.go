package dbsource

import (
	"fmt"
	"go-learn/xorm/utils"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	master = &DbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
	}
	slave = &DbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
	}
)

func init() {
	utils.ViperConf("conf/db-slave.json", master)
	utils.ViperConf("conf/db-master.json", master)
}

type DbSource struct {
	Engine      *xorm.Engine
	Lock        *sync.Mutex
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Host        string `json:"Host"`
	Port        string `json:"Port"`
	Database    string `json:"Database"`
	MaxIdleConn int    `json:"MaxIdleConn"`
	MaxOpenConn int    `json:"MaxOpenConn"`
}

func Slave() *xorm.Engine {
	if slave.Engine != nil {
		return slave.Engine
	}
	slave.Lock.Lock()
	defer slave.Lock.Unlock()

	if slave.Engine != nil {
		return slave.Engine
	}
	engine := createEngine(slave)
	slave.Engine = engine
	return engine
}

func Master() *xorm.Engine {
	if master.Engine != nil {
		return master.Engine
	}
	master.Lock.Lock()
	defer master.Lock.Unlock()
	if master.Engine != nil {
		return master.Engine
	}
	engine := createEngine(master)
	master.Engine = engine
	return engine
}

func CloseEngine() {
	err := master.Engine.Close()
	if err != nil {
		panic(err)
	}
	err = slave.Engine.Close()
	if err != nil {
		panic(err)
	}
}

func createEngine(source *DbSource) *xorm.Engine {
	driver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		source.Username,
		source.Password,
		source.Host,
		source.Port,
		source.Database,
	)
	engine, err := xorm.NewEngine("mysql", driver)
	if err != nil {
		log.Fatalf("Failed to create database engine: %v", err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(source.MaxIdleConn) //最大空闲
	engine.SetMaxOpenConns(source.MaxOpenConn) //最大连接
	return engine
}
