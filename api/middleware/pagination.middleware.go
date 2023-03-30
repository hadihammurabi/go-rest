package middleware

import (
	"encoding/json"
	"go-rest/api/dto"

	"github.com/gofiber/fiber/v2"
)

// Pagination func
func (m Middlewares) Pagination(c *fiber.Ctx) error {
	var pagination dto.PaginationReq
	if err := c.QueryParser(&pagination); err != nil {
		pagination = dto.NewPaginationReq(dto.PaginationReq{})
	} else {
		pagination = dto.NewPaginationReq(pagination)
	}

	filterQ := c.Query("filter", "{}")
	err := json.Unmarshal([]byte(filterQ), &pagination.Filter)
	if err != nil {
		pagination.Filter = make(map[string]string)
	}

	sortQ := c.Query("sort", "{}")
	err = json.Unmarshal([]byte(sortQ), &pagination.Sort)
	if err != nil {
		pagination.Sort = make(map[string]string)
	}

	c.Locals("pagination", pagination)

	return c.Next()
}
