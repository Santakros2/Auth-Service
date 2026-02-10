package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret string
}

func NewService(s string) *Service {
	return &Service{
		secret: s,
	}
}

func (s *Service) LoginService(ctx context.Context, email, password string) (string, error) {

	// DB check now fake validation
	if email != "user@test.com" || password != "1234" {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte(s.secret))

}
