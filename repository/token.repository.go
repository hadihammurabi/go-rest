package repository

import (
	"context"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/ioc"
	"github.com/gowok/qry"
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
	q := qry.Insert(table.NameTokens).
		Column("user_id", "token", "expired_at").
		Values("?", "?", "?")

	_, err := r.db.Exec(q.SQL(), token.UserID, token.Token, token.ExpiredAt)
	if err != nil {
		return token, err
	}

	return token, nil
}

// FindByUserID func
func (r TokenSQL) FindByUserID(c context.Context, id string) (*entity.Token, error) {
	q := qry.Select("user_id", "token", "expired_at").
		From(table.NameTokens).
		Where("user_id = ?")
	row := r.db.QueryRow(q.SQL(), id)

	tokenTable := table.Token{}
	err := row.Scan(
		&tokenTable.UserID,
		&tokenTable.Token,
		&tokenTable.ExpiredAt,
	)

	return tokenTable.ToEntity(), err
}

// FindByToken func
func (r TokenSQL) FindByToken(c context.Context, token string) (*entity.Token, error) {
	q := qry.Select("user_id", "token", "expired_at").
		From(table.NameTokens).
		Where("token = ?")
	row := r.db.QueryRow(q.SQL(), token)

	tokenTable := table.Token{}
	err := row.Scan(
		&tokenTable.UserID,
		&tokenTable.Token,
		&tokenTable.ExpiredAt,
	)

	return tokenTable.ToEntity(), err
}
