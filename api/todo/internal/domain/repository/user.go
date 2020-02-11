package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// UserRepository - UserRepositoryインターフェース
type UserRepository interface {
	Authentication(ctx context.Context) (*domain.User, error)
	GroupIDExistsInGroupIDs(ctx context.Context, groupID string, u *domain.User) bool
}
