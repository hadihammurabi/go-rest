package api

import (
	"go-rest/api/middleware"
	"go-rest/driver"
	"go-rest/service"

	"github.com/gowok/ioc"
	"golang.org/x/exp/slices"

	"github.com/gofiber/fiber/v2"
)

type Policy struct {
	Middlewares *middleware.Middlewares
	AuthService AuthService
	rbac        *driver.RBAC
}

// NewPolicy func
func NewPolicy() *fiber.App {
	api := Policy{
		Middlewares: ioc.MustGet(middleware.Middlewares{}),
		AuthService: ioc.MustGet(service.AuthService{}),
		rbac:        ioc.MustGet(driver.RBAC{}),
	}

	router := fiber.New()
	router.Get("roles",
		api.Middlewares.AuthBearer,
		api.Middlewares.RBAC("policies", driver.RBACRead),
		api.Middlewares.Pagination,
		api.GetAllRoles,
	)
	router.Post("",
		api.Middlewares.AuthBearer,
		api.Middlewares.RBAC("policies", driver.RBACCreate),
		api.AddPolicy,
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
	return Ok(c, api.rbac.GetAllRoles())
}

// // Login func
func (api Policy) AddPolicy(c *fiber.Ctx) error {
	input := make([]any, 0)
	if err := c.BodyParser(&input); err != nil {
		return Fail(c, err)
	}

	if len(input) < 3 {
		return Fail(c, "invalid policy data")
	}

	if !slices.Contains(driver.RBACAll, input[2].(string)) {
		return Fail(c, "invalid policy data")
	}

	_, err := api.rbac.AddPolicy(input...)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, "success")
}

// DeleteRole func
func (api Policy) DeleteRole(c *fiber.Ctx) error {
	name := c.Params("name")
	_, err := api.rbac.DeleteRole(name)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, "success")
}
