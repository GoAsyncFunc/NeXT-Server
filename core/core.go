package core

import (
	"github.com/GoAsyncFunc/NeXT-Server/conf"
)

var newCoreFunc func(c *conf.XrayConfig, log *conf.LogConfig) (Core, error)

func NewCore(c *conf.XrayConfig, log *conf.LogConfig) (Core, error) {
	if newCoreFunc == nil {
		panic("xray core not registered, check build tags")
	}
	return newCoreFunc(c, log)
}

func RegisterCore(f func(c *conf.XrayConfig, log *conf.LogConfig) (Core, error)) {
	newCoreFunc = f
}
