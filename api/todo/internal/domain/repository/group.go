package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// GroupRepository - GroupRepositoryインターフェース
type GroupRepository interface {
	Show(ctx context.Context, groupID string) (*domain.Group, error)
}
