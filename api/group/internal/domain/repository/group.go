package repository

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
)

// GroupRepository - GroupRepositoryインターフェース
type GroupRepository interface {
	Index(ctx context.Context, u *domain.User) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, g *domain.Group) error
	Update(ctx context.Context, g *domain.Group) error
}
