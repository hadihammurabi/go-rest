package repository

import (
	"context"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/qry"
	"gorm.io/gorm"
)

// TokenSQL struct
type TokenSQL struct {
	db *gorm.DB
}

// NewToken func
func NewToken() *TokenSQL {
	return &TokenSQL{
		db: driver.GetSQL(),
	}
}

// Create func
func (r TokenSQL) Create(c context.Context, token *entity.Token) (*entity.Token, error) {
	q := qry.Insert(table.NameTokens)
	q.Column("user_id", "token", "expired_at")
	q.Values("?", "?", "?")

	err := r.db.Exec(q.SQL(), token.UserID, token.Token, token.ExpiredAt).Error
	if err != nil {
		return token, err
	}

	return token, nil
}

// FindByUserID func
func (r TokenSQL) FindByUserID(c context.Context, id string) (*entity.Token, error) {
	tokenTable := table.Token{}
	q := qry.Select("user_id", "token", "expired_at").
		From(table.NameTokens).
		Where("user_id = ?")
	err := r.db.Raw(q.SQL(), id).Scan(&tokenTable).Error
	return tokenTable.ToEntity(), err
}

// FindByToken func
func (r TokenSQL) FindByToken(c context.Context, token string) (*entity.Token, error) {
	tokenTable := table.Token{}
	q := qry.Select("user_id", "token", "expired_at").
		From(table.NameTokens).
		Where("token = ?")
	err := r.db.Raw(q.SQL(), token).Scan(&tokenTable).Error
	return tokenTable.ToEntity(), err
}
