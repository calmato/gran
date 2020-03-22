package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/domain/validation"
)

// TaskService - TaskServiceインターフェース
type TaskService interface {
	Create(ctx context.Context, boardID string, t *domain.Task) (*domain.Task, error)
}

type taskService struct {
	taskDomainValidation validation.TaskDomainValidation
	taskRepository       repository.TaskRepository
}

// NewTaskService - TaskServiceの生成
func NewTaskService(tdv validation.TaskDomainValidation, tr repository.TaskRepository) TaskService {
	return &taskService{
		taskDomainValidation: tdv,
		taskRepository:       tr,
	}
}

func (ts *taskService) Create(
	ctx context.Context, boardID string, t *domain.Task,
) (*domain.Task, error) {
	if ves := ts.taskDomainValidation.Task(ctx, t); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()

	t.ID = uuid.New().String()
	t.CreatedAt = current
	t.UpdatedAt = current

	if err := ts.taskRepository.Create(ctx, t); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return t, nil
}
