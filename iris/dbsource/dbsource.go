package dbsource

import (
	"fmt"
	"go-learn/iris/config"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	master = &dbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
		DbCfg:  config.DB.Master,
	}
	slave = &dbSource{
		Engine: nil,
		Lock:   &sync.Mutex{},
		DbCfg:  config.DB.Slave,
	}
)

type dbSource struct {
	Engine *xorm.Engine
	Lock   *sync.Mutex
	DbCfg  *config.DbCfg
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

func createEngine(source *dbSource) *xorm.Engine {
	cfg := source.DbCfg
	driver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	engine, err := xorm.NewEngine("mysql", driver)
	if err != nil {
		log.Fatalf("Failed to create database engine: %v", err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(cfg.MaxIdleConn) //最大空闲
	engine.SetMaxOpenConns(cfg.MaxOpenConn) //最大连接
	return engine
}
