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
	user := NewUser()
	token := NewToken()

	ioc.Set(func() UserSQL { return *user })
	ioc.Set(func() TokenSQL { return *token })

	repo := Repository{
		User:  user,
		Token: token,
	}

	ioc.Set(func() Repository {
		return repo
	})

	return repo
}
