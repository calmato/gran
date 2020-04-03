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
	Update(ctx context.Context, b *domain.Board) error
	IndexBoardList(ctx context.Context, groupID string, boardID string) ([]*domain.BoardList, error)
	ShowBoardList(ctx context.Context, groupID string, boardID string, boardListID string) (*domain.BoardList, error)
	CreateBoardList(ctx context.Context, groupID string, boardID string, bl *domain.BoardList) error
	UpdateBoardList(ctx context.Context, groupID string, boardID string, bl *domain.BoardList) error
}
