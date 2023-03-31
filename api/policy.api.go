package api

import (
	"go-rest/api/middleware"
	"go-rest/driver"
	"go-rest/service"

	"github.com/gowok/ioc"

	"github.com/gofiber/fiber/v2"
)

type Policy struct {
	Middlewares   *middleware.Middlewares
	AuthService   AuthService
	PolicyService PolicyService
}

// NewPolicy func
func NewPolicy() *fiber.App {
	api := Policy{
		Middlewares:   ioc.MustGet(middleware.Middlewares{}),
		AuthService:   ioc.MustGet(service.AuthService{}),
		PolicyService: ioc.MustGet(service.PolicyService{}),
	}

	router := fiber.New()
	router.Post("",
		api.Middlewares.AuthBearer,
		api.Middlewares.RBAC("policies", driver.RBACCreate),
		api.AddPolicy,
	)
	router.Get("roles",
		api.Middlewares.AuthBearer,
		api.Middlewares.RBAC("policies", driver.RBACRead),
		api.Middlewares.Pagination,
		api.GetAllRoles,
	)
	router.Delete("roles/:name",
		api.Middlewares.AuthBearer,
		api.Middlewares.RBAC("policies", driver.RBACDelete),
		api.DeleteRole,
	)

	return router
}

// GetAllRoles func
func (api Policy) GetAllRoles(c *fiber.Ctx) error {
	return Ok(c, api.PolicyService.GetAllRoles(c.Context()))
}

// // Login func
func (api Policy) AddPolicy(c *fiber.Ctx) error {
	input := make([]any, 0)
	if err := c.BodyParser(&input); err != nil {
		return Fail(c, err)
	}

	err := api.PolicyService.AddPolicy(c.Context(), input)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, "success")
}

// DeleteRole func
func (api Policy) DeleteRole(c *fiber.Ctx) error {
	name := c.Params("name")
	err := api.PolicyService.DeleteRole(c.Context(), name)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, "success")
}
