package repository

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
)

// UserRepository - UserRepositoryインターフェース
type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
