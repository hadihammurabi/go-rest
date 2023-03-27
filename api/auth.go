package api

import (
	"go-rest/api/dto"
	"go-rest/api/middleware"
	"go-rest/service"

	"go-rest/entity"

	"github.com/gowok/ioc"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Middlewares *middleware.Middlewares
	AuthService AuthService
}

// NewAuth func
func NewAuth() *fiber.App {
	api := Auth{
		Middlewares: ioc.MustGet(middleware.Middlewares{}),
		AuthService: ioc.MustGet(service.AuthService{}),
	}

	router := fiber.New()
	router.Post("/login", api.Login)
	router.Get("/me", api.Middlewares.AuthBearer, api.Me)

	return router
}

// Login func
func (api Auth) Login(c *fiber.Ctx) error {
	input := &dto.UserLoginRequest{}
	if err := c.BodyParser(input); err != nil {
		return Fail(c, err)
	}

	user := &entity.User{
		Email:    input.Email,
		Password: input.Password,
	}

	token, err := api.AuthService.Login(c.Context(), user)
	if err != nil {
		return Fail(c, "invalid credentials")
	}

	return Ok(c, &dto.UserLoginResponse{
		Token: token,
		Type:  "Bearer",
	})
}

// Me func
func (api Auth) Me(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*entity.User)
	return Ok(c, *fromLocals)
}
