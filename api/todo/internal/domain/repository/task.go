package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// TaskRepository - TaskRepositoryインターフェース
type TaskRepository interface {
	IndexByBoardID(ctx context.Context, boardID string) ([]*domain.Task, error)
	Create(ctx context.Context, t *domain.Task) error
}
