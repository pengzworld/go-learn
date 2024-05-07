package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"go-learn/gin/utils"
)

var C = struct {
	App struct {
		Name        string       `yaml:"Name"`
		IP          string       `yaml:"IP"`
		Host        string       `yaml:"Host"`
		Port        int          `yaml:"Port"`
		TimeFormat  string       `yaml:"TimeFormat"`
		LogLevel    logrus.Level `yaml:"LogLevel"`
		ReleaseMode string       `yaml:"ReleaseMode"`
	}
}{
	// other default values...
}

func InitConfig() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	file := fmt.Sprintf("config/%s.yml", env)
	utils.ViperConf(file, &C)
}
