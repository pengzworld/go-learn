package config

import (
	"fmt"
	"go-learn/iris/utils"
	"os"

	"github.com/kataras/iris/v12"
)

func init() {
	loadConfiguration()
}

var C = struct {
	Iris iris.Configuration
	App  struct {
		Name       string
		IP         string
		Host       string
		Port       int
		TimeFormat string
	}
}{
	Iris: iris.DefaultConfiguration(),
	// other default values...
}

type DbCfg struct {
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Host        string `json:"Host"`
	Port        string `json:"Port"`
	Database    string `json:"Database"`
	MaxIdleConn int    `json:"MaxIdleConn"`
	MaxOpenConn int    `json:"MaxOpenConn"`
}

var DB = struct {
	Master *DbCfg
	Slave  *DbCfg
}{}

func loadConfiguration() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	file := fmt.Sprintf("config/%s.yml", env)
	utils.ViperConf(file, &C)
	utils.ViperConf("config/db.json", &DB)
}
