package repository

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
)

// UserRepository - UserRepositoryインターフェース
type UserRepository interface {
	Authentication(ctx context.Context) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
}
