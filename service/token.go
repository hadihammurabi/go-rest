package service

import (
	"context"
	"go-rest/repository"

	"go-rest/entity"

	"github.com/gowok/ioc"
)

// TokenService struct
type TokenService struct {
	TokenRepo TokenRepo
}

// NewTokenService func
func NewTokenService() TokenService {
	return TokenService{
		TokenRepo: ioc.MustGet(repository.TokenSQL{}),
	}
}

// Create func
func (u TokenService) Create(c context.Context, token *entity.Token) (*entity.Token, error) {
	tokenFromTable, err := u.TokenRepo.Create(c, token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByUserID func
func (u TokenService) FindByUserID(c context.Context, id string) (*entity.Token, error) {
	tokenFromTable, err := u.TokenRepo.FindByUserID(c, id)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByToken func
func (u TokenService) FindByToken(c context.Context, token string) (*entity.Token, error) {
	tokenFromTable, err := u.TokenRepo.FindByToken(c, token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}
