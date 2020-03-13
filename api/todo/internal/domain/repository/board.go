package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// BoardRepository - BoardRepositoryインターフェース
type BoardRepository interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
	Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error)
	Create(ctx context.Context, b *domain.Board) error
	IndexBoardList(ctx context.Context, groupID string, boardID string) ([]*domain.BoardList, error)
}
