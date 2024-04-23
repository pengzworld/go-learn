package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
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

func loadConfiguration() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	configFile := fmt.Sprintf("config/%s.yml", env)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := viper.Unmarshal(&C); err != nil {
			panic(err)
		}
		fmt.Println(&C)
	})
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
