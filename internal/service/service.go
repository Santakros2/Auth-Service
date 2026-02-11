package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"server.com/auth-service/internal/domain"
	"server.com/auth-service/internal/repository"
	"server.com/auth-service/pkg/hash"
)

type Service struct {
	secret string
	repo   repository.UserRepository
}

func NewService(s string, r repository.UserRepository) *Service {
	return &Service{
		secret: s,
		repo:   r,
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

type SignupResult struct {
	Token string
	User  *domain.User
}

func (s *Service) SignupService(ctx context.Context, name, email, password string) (*SignupResult, error) {

	// check if user exists
	exists, err := s.repo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("user already exists")
	}

	// hash password
	hashPass, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// create domain user
	user := domain.NewUser(email, name, hashPass)

	// save user
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	return &SignupResult{
		Token: signedToken,
		User:  user,
	}, nil
}
