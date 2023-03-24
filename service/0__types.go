package service

import (
	"context"

	"go-rest/api/dto"
	"go-rest/entity"
)

type UserRepo interface {
	All(c context.Context, pagination dto.PaginationReq) (dto.PaginationRes[entity.User], error)
	ChangePassword(c context.Context, id string, password string) (*entity.User, error)
	Create(c context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	FindByID(c context.Context, id string) (*entity.User, error)
}

type TokenRepo interface {
	Create(c context.Context, token *entity.Token) (*entity.Token, error)
	FindByToken(c context.Context, token string) (*entity.Token, error)
	FindByUserID(c context.Context, id string) (*entity.Token, error)
}
