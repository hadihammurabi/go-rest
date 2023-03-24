package api

import (
	"go-rest/api/middleware"
	"go-rest/driver"
	"go-rest/service"
	"log"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"

	"github.com/gofiber/fiber/v2"
)

// ConfigureRoute do HTTP routing.
// Here, we can set route or just invoke another func to do modular route.
func (api *APIRest) ConfigureRoute() {
	api.HTTP.Mount("/", NewIndex())
	api.HTTP.Mount("/users", NewUser())
}

// APIRest struct
type APIRest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
	Service     *service.Service
	Validator   *gowok.Validator
	Config      *gowok.Config
}

func NewAPIRest() *APIRest {
	api := &APIRest{
		HTTP:      driver.NewAPI(),
		Config:    ioc.Get(gowok.Config{}),
		Validator: ioc.Get(gowok.Validator{}),
		Service:   service.New(),
	}
	api.ConfigureMiddleware()
	api.ConfigureRoute()
	return api
}

// ConfigureMiddleware func
func (api *APIRest) ConfigureMiddleware() {
	api.Middlewares = middleware.NewMiddleware(
		api.Config,
		api.Service,
	)
}

func (d *APIRest) Run() {
	log.Println("API REST started at", d.Config.App.Rest.Host)
	if err := d.HTTP.Listen(d.Config.App.Rest.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *APIRest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}

func Run() {
	NewAPIRest().Run()
}
