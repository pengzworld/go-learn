package dbsource

import (
	"fmt"
	"go-learn/iris/config"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DbSource struct {
	Engine *xorm.Engine
	Lock   *sync.Mutex
	DbCfg  *config.DbCfg
}

var (
	master = &DbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
		DbCfg:  config.DB.Master,
	}
	slave = &DbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
		DbCfg:  config.DB.Slave,
	}
)

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
		source.DbCfg.Username,
		source.DbCfg.Password,
		source.DbCfg.Host,
		source.DbCfg.Port,
		source.DbCfg.Database,
	)
	engine, err := xorm.NewEngine("mysql", driver)
	if err != nil {
		log.Fatalf("Failed to create database engine: %v", err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(source.DbCfg.MaxIdleConn) //最大空闲
	engine.SetMaxOpenConns(source.DbCfg.MaxOpenConn) //最大连接
	return engine
}
