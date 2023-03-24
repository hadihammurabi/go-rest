package service

import "github.com/gowok/ioc"

type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
}

func New() *Service {
	auth := NewAuthService()
	user := NewUserService()
	token := NewTokenService()

	ioc.Set(func() AuthService { return auth })
	ioc.Set(func() UserService { return user })
	ioc.Set(func() TokenService { return token })

	return &Service{
		Auth:  auth,
		User:  user,
		Token: token,
	}
}
