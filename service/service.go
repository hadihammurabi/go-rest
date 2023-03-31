package service

import "github.com/gowok/ioc"

type Service struct {
	Auth   AuthService
	User   UserService
	Token  TokenService
	Policy PolicyService
}

func New() *Service {
	return &Service{
		Auth:   NewAuthService(),
		User:   NewUserService(),
		Token:  NewTokenService(),
		Policy: NewPolicyService(),
	}
}

func Init() {
	sv := New()

	ioc.Set(func() Service { return *sv })
	ioc.Set(func() AuthService { return sv.Auth })
	ioc.Set(func() UserService { return sv.User })
	ioc.Set(func() TokenService { return sv.Token })
	ioc.Set(func() PolicyService { return sv.Policy })
}
