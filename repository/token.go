package repository

import (
	"context"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/ioc"
)

// TokenSQL struct
type TokenSQL struct {
	db *driver.DB
}

// NewToken func
func NewToken() *TokenSQL {
	return &TokenSQL{
		db: ioc.Get(driver.DB{}),
	}
}

// Create func
func (r TokenSQL) Create(c context.Context, token *entity.Token) (*entity.Token, error) {
	tokenTable := table.TokenFromEntity(token)
	// tokenTable.ID = uuid.NewString()
	_, err := r.db.NewInsert().
		Model(tokenTable).
		Exec(c)
	if err != nil {
		return token, err
	}

	return tokenTable.ToEntity(), nil
}

// FindByUserID func
func (r TokenSQL) FindByUserID(c context.Context, id string) (*entity.Token, error) {
	tokenTable := &table.Token{}
	err := r.db.NewSelect().Model(&tokenTable).Where("id", id).Scan(c)
	return tokenTable.ToEntity(), err
}

// FindByToken func
func (r TokenSQL) FindByToken(c context.Context, token string) (*entity.Token, error) {
	tokenTable := &table.Token{}
	err := r.db.NewSelect().Model(&tokenTable).Where("token", token).Scan(c)
	return tokenTable.ToEntity(), err
}
