package repository

import (
	"context"

	"server.com/auth-service/internal/domain"
)

type UserRepo struct {
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

func NewRepo() UserRepository {
	return &UserRepo{}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return false, nil
}
