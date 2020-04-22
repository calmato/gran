package repository

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
)

// BoardRepository - BoardRepositoryインターフェース
type BoardRepository interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
}
