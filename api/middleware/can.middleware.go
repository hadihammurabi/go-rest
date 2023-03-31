package middleware

import (
	"go-rest/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Can func
func (m Middlewares) Can(act, obj string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userFromLocals := c.Locals("user")
		if userFromLocals == nil {
			return c.Status(http.StatusUnauthorized).SendString("unauthorized")
		}

		user := userFromLocals.(*entity.User)

		var allowed bool
		var err error
		if user.Email == "root" {
			allowed, err = m.pol.Enforce(user.Email, obj, act)
		} else {
			allowed, err = m.pol.Enforce(user.ID, obj, act)
		}

		if err != nil || !allowed {
			return c.Status(http.StatusUnauthorized).SendString("unauthorized")
		}

		return c.Next()
	}
}
