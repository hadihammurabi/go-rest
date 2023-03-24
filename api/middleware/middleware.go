package middleware

import (
	"go-rest/service"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

// Middlewares type
type Middlewares struct {
	config  *gowok.Config
	service *service.Service
}

func NewMiddleware(config *gowok.Config, service *service.Service) Middlewares {
	middlewares := Middlewares{
		config:  config,
		service: service,
	}

	ioc.Set(func() Middlewares { return middlewares })

	return middlewares
}