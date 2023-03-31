package service

import (
	"context"
	"database/sql"
	"errors"
	"go-rest/driver"
	"go-rest/repository"

	"go-rest/api/dto"
	"go-rest/entity"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/gowok/ioc"
)

// UserService struct
type UserService struct {
	UserRepo UserRepo
	config   *gowok.Config
	rbac     *driver.RBAC
}

// NewUserService func
func NewUserService() UserService {
	return UserService{
		UserRepo: ioc.MustGet(repository.UserSQL{}),
		config:   ioc.MustGet(gowok.Config{}),
		rbac:     ioc.MustGet(driver.RBAC{}),
	}
}

// All func
func (u UserService) All(c context.Context, pagination dto.PaginationReq) (res dto.PaginationRes[entity.User], err error) {
	res, err = u.UserRepo.All(c, pagination)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Create func
func (u UserService) Create(c context.Context, user *entity.User) (*entity.User, error) {
	isEmailExists, err := u.IsEmailUsed(c, user.Email)
	if err != nil {
		return nil, err
	}

	if isEmailExists {
		return nil, errors.New("email already used")
	}

	pass := hash.PasswordHash(user.Password, u.config.App.Key)
	user.Password = pass.Hashed

	userFromTable, err := u.UserRepo.Create(c, user)
	if err != nil {
		return nil, err
	}

	u.rbac.AddGroupingPolicy(user.ID, "staff")

	return userFromTable, nil
}

func (u UserService) IsEmailUsed(c context.Context, email string) (bool, error) {
	_, err := u.FindByEmail(c, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

// FindByEmail func
func (u UserService) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userFromTable, err := u.UserRepo.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}

// FindByID func
func (u UserService) FindByID(c context.Context, id string) (*entity.User, error) {
	userFromTable, err := u.UserRepo.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}

// ChangePassword func
func (u UserService) ChangePassword(c context.Context, id string, password string) (*entity.User, error) {
	pass := hash.PasswordHash(password, u.config.App.Key)
	userFromTable, err := u.UserRepo.ChangePassword(c, id, pass.Hashed)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}
