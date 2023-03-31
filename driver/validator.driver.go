package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

func initValidator() {
	v := gowok.NewValidator()
	ioc.Set(func() gowok.Validator { return *v })
}
