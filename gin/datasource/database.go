package datasource

import (
	"fmt"
	"go-learn/gin/core"
	"sync"

	"go-learn/gin/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Database *engine
var configs = make(map[string]*conf)

type engine struct {
	sync.RWMutex
	engine map[string]*xorm.Engine
}

type conf struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    string `json:"database"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxOpenConn int    `json:"maxOpenConn"`
}

func InitDatabase() {
	utils.ViperConf("config/db.json", &configs)
	Database = &engine{
		engine: make(map[string]*xorm.Engine),
	}
	for k, c := range configs {
		e, ok := Database.Get(k)
		if !ok {
			e = create(c)
			Database.save(k, e)
		}
		if err := e.Ping(); err != nil {
			core.Logger.Printf("Failed to ping database %s: %v", k, err)
			continue
		}
	}
}

func (e *engine) Get(key string) (*xorm.Engine, bool) {
	e.RLock()
	val, ok := e.engine[key]
	e.RUnlock()
	return val, ok
}

func (e *engine) save(key string, value *xorm.Engine) {
	e.Lock()
	e.engine[key] = value
	e.Unlock()
}

func (e *engine) Close() {
	for k, v := range e.engine {
		if err := v.Close(); err != nil {
			core.Logger.Printf("Failed to close database engine:%s: %v", k, err)
		}
	}
}

func create(c *conf) *xorm.Engine {
	driver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
	engine, err := xorm.NewEngine("mysql", driver)
	if err != nil {
		core.Logger.Printf("Failed to create database engine: %v", err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(c.MaxIdleConn) //最大空闲
	engine.SetMaxOpenConns(c.MaxOpenConn) //最大连接
	return engine
}
