package repository

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
)

// TaskRepository - TaskRepositoryインターフェース
type TaskRepository interface {
	IndexByBoardID(ctx context.Context, boardID string) ([]*domain.Task, error)
	Show(ctx context.Context, taskID string) (*domain.Task, error)
	Create(ctx context.Context, t *domain.Task) error
}
