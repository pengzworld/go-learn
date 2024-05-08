package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"go-learn/gin/utils"
)

var C = struct {
	App struct {
		Name        string `yaml:"Name"`
		IP          string `yaml:"IP"`
		Host        string `yaml:"Host"`
		Port        int    `yaml:"Port"`
		TimeFormat  string `yaml:"TimeFormat"`
		ReleaseMode string `yaml:"ReleaseMode"`
		Log         struct {
			Path              string       `yaml:"Path"`
			MaxSize           int          `yaml:"MaxSize"`
			MaxBackups        int          `yaml:"MaxBackups"`
			MaxAge            int          `yaml:"MaxAge"`
			Compress          bool         `yaml:"Compress"`
			ReportCaller      bool         `yaml:"ReportCaller"`
			Level             logrus.Level `yaml:"Level"`
			PrettyPrint       bool         `yaml:"PrettyPrint"`
			DisableHTMLEscape bool         `yaml:"DisableHTMLEscape"`
		}
		AccessLog struct {
			Path              string       `yaml:"Path"`
			MaxSize           int          `yaml:"MaxSize"`
			MaxBackups        int          `yaml:"MaxBackups"`
			MaxAge            int          `yaml:"MaxAge"`
			Compress          bool         `yaml:"Compress"`
			ReportCaller      bool         `yaml:"ReportCaller"`
			Level             logrus.Level `yaml:"Level"`
			PrettyPrint       bool         `yaml:"PrettyPrint"`
			DisableHTMLEscape bool         `yaml:"DisableHTMLEscape"`
		}
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
