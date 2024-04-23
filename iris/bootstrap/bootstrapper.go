package bootstrap

import (
	"go-learn/iris/config"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
}

// New returns a new Bootstrapper.
func New(cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func SetLogger(b *Bootstrapper) {
	b.Logger().SetLevel(config.C.Iris.LogLevel)
	b.Logger().SetOutput(makeLog())
	b.Logger().AddOutput(os.Stdout)
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.HandleDir("/public", iris.Dir("./public/"))
	return b
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func makeLog() *rotatelogs.RotateLogs {
	log := "./runtimes/logs/log.%Y-%m-%d-%H"
	w, err := rotatelogs.New(
		log,
		rotatelogs.WithMaxAge(24*time.Hour),    //最大保存时间
		rotatelogs.WithRotationTime(time.Hour)) //文件的轮转时间间隔
	if err != nil {
		panic(err)
	}
	return w
}
