package core

import (
	"strings"
	"sync"
)

var sMap sync.Map
var Container *container

type container struct {
}

func InitContainer() {
	Container = NewContainer()
}

func NewContainer() *container {
	return &container{}
}

func (c *container) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

func (c *container) Set(key string, value interface{}) (res bool) {
	if _, exists := c.KeyIsExists(key); exists == false {
		sMap.Store(key, value)
		res = true
	}
	return
}

func (c *container) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

func (c *container) Delete(key string) {
	sMap.Delete(key)
}

func (c *container) FuzzyDelete(keyPre string) {
	sMap.Range(func(key, value interface{}) bool {
		if val, ok := key.(string); ok {
			if strings.HasPrefix(val, keyPre) {
				sMap.Delete(val)
			}
		}
		return true
	})
}
