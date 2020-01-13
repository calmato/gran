package repository

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
)

// GroupRepository - GroupRepositoryインターフェース
type GroupRepository interface {
	Create(ctx context.Context, g *domain.Group) error
}
