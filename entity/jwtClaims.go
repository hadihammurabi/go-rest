package entity

import (
	"github.com/golang-jwt/jwt"
)

// JWTClaims model
type JWTClaims struct {
	*jwt.StandardClaims
	UserID string `json:"user_id,omitempty"`
}
