package utils

import (
	"fmt"
	"os"
	"path/filepath"

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

// CreateFileIfNotExists checks if a file exists and creates it if not.
// It also creates the directory if it doesn't exist.
func CreateFileIfNotExists(filePath string) error {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, need to create
		// Get directory path
		dirPath := filepath.Dir(filePath)
		// Check if directory exists
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// Directory doesn't exist, create recursively
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			fmt.Printf("Directory created: %s\n", dirPath)
		}
		// Create file
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()
		fmt.Printf("File created: %s\n", filePath)
	} else if err != nil {
		// Other error
		return fmt.Errorf("error checking file status: %v", err)
	} else {
		fmt.Printf("File already exists: %s\n", filePath)
	}

	return nil
}
