package middleware

import (
	"go-rest/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// RBAC func
func (m Middlewares) RBAC(obj, act string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userFromLocals := c.Locals("user")
		if userFromLocals == nil {
			return c.Status(http.StatusUnauthorized).SendString("unauthorized")
		}

		user := userFromLocals.(*entity.User)

		var allowed bool
		var err error
		if user.Email == "root" {
			allowed, err = m.rbac.Enforce(user.Email, obj, act)
		} else {
			allowed, err = m.rbac.Enforce(user.ID, obj, act)
		}

		if err != nil || !allowed {
			return c.Status(http.StatusUnauthorized).SendString("unauthorized")
		}

		return c.Next()
	}
}
