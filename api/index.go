package api

import (
	"github.com/gofiber/fiber/v2"
)

type Index struct{}

func NewIndex() *fiber.App {
	api := Index{}

	router := fiber.New()
	router.Get("", api.Index)

	return router
}

func (api Index) Index(c *fiber.Ctx) error {
	return Ok(c, fiber.Map{
		"message": "Selamat datang di Go REST API",
	})
}
