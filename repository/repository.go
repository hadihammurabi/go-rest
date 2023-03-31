package repository

import (
	"github.com/gowok/ioc"
)

// Repository struct
type Repository struct {
	User  *UserSQL
	Token *TokenSQL
}

// NewRepository func
func NewRepository() Repository {
	repo := Repository{
		User:  NewUser(),
		Token: NewToken(),
	}

	return repo
}

func Init() {
	repo := NewRepository()
	ioc.Set(func() UserSQL { return *repo.User })
	ioc.Set(func() TokenSQL { return *repo.Token })
	ioc.Set(func() Repository {
		return repo
	})
}
