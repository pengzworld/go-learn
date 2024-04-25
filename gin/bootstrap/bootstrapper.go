package bootstrap

import "github.com/gin-gonic/gin"

type Configurator func(b *gin.Engine)

func New(cfgs ...Configurator) *gin.Engine {
	b := gin.New()
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}
