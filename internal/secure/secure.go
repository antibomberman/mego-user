package secure

import (
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Secure interface {
	Generate(userId string) (string, error)
	Check(tokenString string) (bool, error)
	Parse(tokenString string) (*Claims, error)
}

type secure struct {
	cfg *config.Config
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewSecure(cfg *config.Config) Secure {
	return &secure{
		cfg: cfg,
	}
}
func (a secure) Generate(UserId string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.cfg.JWTSecret))
}
func (a secure) Check(tokenString string) (bool, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JWTSecret), nil
	})
	return err == nil && token.Valid, err
}
func (a secure) Parse(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JWTSecret), nil
	})
	return claims, err
}
