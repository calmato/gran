package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/repository"
	"github.com/calmato/gran/api/todo/internal/domain/validation"
)

// TaskService - TaskServiceインターフェース
type TaskService interface {
	Show(ctx context.Context, taskID string) (*domain.Task, error)
	Create(ctx context.Context, groupID string, boardID string, boardListID string, t *domain.Task) (*domain.Task, error)
}

type taskService struct {
	taskDomainValidation validation.TaskDomainValidation
	taskRepository       repository.TaskRepository
	boardRepository      repository.BoardRepository
}

// NewTaskService - TaskServiceの生成
func NewTaskService(
	tdv validation.TaskDomainValidation, tr repository.TaskRepository, br repository.BoardRepository,
) TaskService {
	return &taskService{
		taskDomainValidation: tdv,
		taskRepository:       tr,
		boardRepository:      br,
	}
}

func (ts *taskService) Show(ctx context.Context, taskID string) (*domain.Task, error) {
	t, err := ts.taskRepository.Show(ctx, taskID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return t, nil
}

func (ts *taskService) Create(
	ctx context.Context, groupID string, boardID string, boardListID string, t *domain.Task,
) (*domain.Task, error) {
	if ves := ts.taskDomainValidation.Task(ctx, t); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	bl, err := ts.boardRepository.ShowBoardList(ctx, groupID, boardID, boardListID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	current := time.Now()

	t.ID = uuid.New().String()
	t.CreatedAt = current
	t.UpdatedAt = current

	if err := ts.taskRepository.Create(ctx, t); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	bl.TaskIDs = append(bl.TaskIDs, t.ID)

	if err := ts.boardRepository.UpdateBoardList(ctx, groupID, boardID, bl); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return t, nil
}
