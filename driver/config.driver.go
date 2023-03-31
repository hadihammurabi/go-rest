package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

func NewConfig() (gowok.Config, error) {
	conf, err := gowok.Configure("config.yaml")
	if err != nil {
		panic(err)
	}

	return conf, err
}

func initConfig() {
	conf, err := NewConfig()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() gowok.Config { return conf })
}
