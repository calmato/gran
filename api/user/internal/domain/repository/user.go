package repository

import (
	"context"

	"github.com/calmato/gran/api/user/internal/domain"
)

// UserRepository - UserRepositoryインターフェース
type UserRepository interface {
	Authentication(ctx context.Context) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
	Update(ctx context.Context, u *domain.User) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
