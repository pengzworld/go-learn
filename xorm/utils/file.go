package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ViperConf(file string, param interface{}) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := viper.Unmarshal(&param); err != nil {
			panic(err)
		}
		fmt.Println(&param)
	})
	if err := viper.Unmarshal(&param); err != nil {
		panic(err)
	}
}
