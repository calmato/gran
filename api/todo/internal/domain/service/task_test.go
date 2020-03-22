package service

import (
	"context"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
	mock_repository "github.com/16francs/gran/api/todo/mock/domain/repository"
	mock_validation "github.com/16francs/gran/api/todo/mock/domain/validation"
	"github.com/golang/mock/gomock"
)

func TestTaskService_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)
	current := time.Now()
	boardID := "task-create-board-id"
	boardListID := "task-create-boardlist-id"

	task := &domain.Task{
		Name:            "タスク",
		Description:     "説明",
		Labels:          []string{},
		AttachmentURLs:  []string{},
		AssignedUserIDs: []string{},
		DeadlinedAt:     current,
	}

	// Defined mocks
	tdvm := mock_validation.NewMockTaskDomainValidation(ctrl)
	tdvm.EXPECT().Task(ctx, task).Return(ves)

	trm := mock_repository.NewMockTaskRepository(ctrl)
	trm.EXPECT().Create(ctx, task).Return(nil)

	// Start test
	target := NewTaskService(tdvm, trm)

	_, err := target.Create(ctx, boardID, boardListID, task)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
