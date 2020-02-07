package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// BoardRepository - BoardRepositoryインターフェース
type BoardRepository interface {
	Create(ctx context.Context, groupID string, b *domain.Board) error
}
