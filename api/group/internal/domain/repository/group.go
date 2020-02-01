package repository

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
)

// GroupRepository - GroupRepositoryインターフェース
type GroupRepository interface {
	Create(ctx context.Context, u *domain.User, g *domain.Group) error
}
