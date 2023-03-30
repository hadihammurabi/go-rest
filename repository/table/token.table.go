package table

import (
	"go-rest/entity"
	"time"
)

const NameTokens string = "auth.tokens"

// Token model
type Token struct {
	Base
	UserID    string     `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// BeforeCreate func
// func (u *Token) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, err := uuid.NewRandom()
// 	u.ID = id
// 	return
// }

// ToEntity func
func (t Token) ToEntity() *entity.Token {
	return &entity.Token{
		UserID:    t.UserID,
		Token:     t.Token,
		ExpiredAt: t.ExpiredAt,
	}
}

// TokenFromEntity func
func TokenFromEntity(e *entity.Token) *Token {
	return &Token{
		UserID:    e.UserID,
		Token:     e.Token,
		ExpiredAt: e.ExpiredAt,
	}
}
