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
	Create(ctx context.Context, groupID string, boardID string, t *domain.Task) (*domain.Task, error)
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

func (ts *taskService) Create(
	ctx context.Context, groupID string, boardID string, t *domain.Task,
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

	bl, err := ts.boardRepository.ShowBoardList(ctx, groupID, boardID, t.BoardListID)
	if err != nil {
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
