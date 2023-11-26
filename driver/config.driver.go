package driver

import (
	"os"

	"github.com/gowok/gowok"
)

func NewConfig() (*gowok.Config, error) {
	conf, err := gowok.NewConfig(os.OpenFile("config.yaml", os.O_RDONLY, 0666))
	if err != nil {
		panic(err)
	}

	return conf, err
}

var conf *gowok.Config

func GetConfig() *gowok.Config {
	if conf != nil {
		return conf
	}

	conf = gowok.Must(NewConfig())
	return conf
}
