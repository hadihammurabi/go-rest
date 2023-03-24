package service

import (
	"context"
	"errors"

	"go-rest/entity"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/gowok/ioc"
)

// AuthService struct
type AuthService struct {
	userService  UserService
	tokenService TokenService
	jwtService   JWTService
	config       *gowok.Config
}

// NewAuthService func
func NewAuthService() AuthService {
	config := ioc.Get(gowok.Config{})
	return AuthService{
		userService:  NewUserService(),
		tokenService: NewTokenService(),
		jwtService:   NewJWTService(),
		config:       config,
	}
}

// Login func
func (a AuthService) Login(c context.Context, userInput *entity.User) (string, error) {
	user, err := a.userService.FindByEmail(c, userInput.Email)
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	isPasswordValid := hash.PasswordVerify(userInput.Password, user.Password, a.config.App.Key)
	if isPasswordValid {
		return "", errors.New("email or password invalid")
	}

	token, err := a.jwtService.Create(user)
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	return token.Token, nil
}
