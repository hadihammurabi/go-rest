package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

func Config() (gowok.Config, error) {
	conf, err := gowok.Configure("config.yaml")
	if err != nil {
		panic(err)
	}

	ioc.Set(func() gowok.Config { return conf })

	return conf, err
}
