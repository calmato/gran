package repository

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// TaskRepository - TaskRepositoryインターフェース
type TaskRepository interface {
	IndexByBoardListID(ctx context.Context, boardListID string) ([]*domain.Task, error)
}
