package repository

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
)

// UserRepository - UserRepositoryインターフェース
type UserRepository interface {
	Authentication(ctx context.Context) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
	CreateGroup(ctx context.Context, u *domain.User) error
}
