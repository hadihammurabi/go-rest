package api

import (
	"context"
	"go-rest/api/dto"

	"go-rest/entity"
)

type AuthService interface {
	Login(c context.Context, userInput *entity.User) (string, error)
}

type UserService interface {
	All(c context.Context, pagination dto.PaginationReq) (res dto.PaginationRes[entity.User], err error)
	Create(c context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	FindByID(c context.Context, id string) (*entity.User, error)
	ChangePassword(c context.Context, id string, password string) (*entity.User, error)
}
