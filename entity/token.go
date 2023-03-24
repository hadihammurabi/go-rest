package entity

import (
	"time"
)

// Token entity
type Token struct {
	UserID    string     `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}
