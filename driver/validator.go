package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

func Validator() {
	v := gowok.NewValidator()
	ioc.Set(func() gowok.Validator { return *v })
}
