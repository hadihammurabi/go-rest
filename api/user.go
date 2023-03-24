package api

import (
	"go-rest/api/dto"
	"go-rest/api/middleware"
	"go-rest/service"

	"go-rest/entity"

	"github.com/gowok/ioc"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Middlewares *middleware.Middlewares
	UserService UserService
}

// NewUser func
func NewUser() *fiber.App {
	api := User{
		Middlewares: ioc.MustGet(middleware.Middlewares{}),
		UserService: ioc.MustGet(service.UserService{}),
	}

	router := fiber.New()
	router.Get("", api.Middlewares.Pagination, api.UserGetAll)
	router.Post("", api.UserCreate)
	router.Put(":id", api.UserUpdateByID)

	return router
}

// UserGetAll func
func (api User) UserGetAll(c *fiber.Ctx) error {
	pagination := c.Locals("pagination").(dto.PaginationReq)

	users, err := api.UserService.All(c.Context(), pagination)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, users)
}

// UserCreate func
func (api User) UserCreate(c *fiber.Ctx) error {
	var in dto.UserCreateReq
	if err := c.BodyParser(&in); err != nil {
		return Fail(c, err)
	}

	user, err := api.UserService.Create(c.Context(), &entity.User{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, user)
}

// UserCreate func
func (api User) UserUpdateByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var in dto.UserUpdateReq
	if err := c.BodyParser(&in); err != nil {
		return Fail(c, err)
	}

	user, err := api.UserService.ChangePassword(c.Context(), id, in.Password)
	if err != nil {
		return Fail(c, err)
	}

	return Ok(c, user)
}
