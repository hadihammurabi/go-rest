package entity

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims model
type JWTClaims struct {
	*jwt.StandardClaims
	UserID string `json:"user_id,omitempty"`
}
