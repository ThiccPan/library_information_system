package config

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredential = errors.New("invalid credential")
)

type AuthJWT struct {
	secret string
}

type JwtCustomClaims struct {
	Id     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	RoleId uint
	jwt.RegisteredClaims
}

func NewAuthJWT(secret string) *AuthJWT {
	return &AuthJWT{
		secret: secret,
	}
}

func (aj *AuthJWT) GetSecret() string {
	return aj.secret
}

func (aj *AuthJWT) GenerateToken(id uint, email string, username string, role_id uint) (string, error) {
	claims := &JwtCustomClaims{
		id,
		email,
		username,
		role_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(aj.secret))
}
