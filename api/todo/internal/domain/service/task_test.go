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
	groupID := "task-create-group-id"
	boardID := "task-create-board-id"
	boardListID := "task-create-boardlist-id"

	task := &domain.Task{
		Name:            "タスク",
		Description:     "説明",
		Labels:          []string{},
		AttachmentURLs:  []string{},
		AssignedUserIDs: []string{},
		BoardListID:     boardListID,
		DeadlinedAt:     current,
	}

	bl := &domain.BoardList{
		ID:      boardListID,
		Name:    "ボードリスト",
		Color:   "",
		TaskIDs: []string{task.ID},
	}

	// Defined mocks
	tdvm := mock_validation.NewMockTaskDomainValidation(ctrl)
	tdvm.EXPECT().Task(ctx, task).Return(ves)

	trm := mock_repository.NewMockTaskRepository(ctrl)
	trm.EXPECT().Create(ctx, task).Return(nil)

	brm := mock_repository.NewMockBoardRepository(ctrl)
	brm.EXPECT().ShowBoardList(ctx, groupID, boardID, boardListID).Return(bl, nil)
	brm.EXPECT().UpdateBoardList(ctx, groupID, boardID, bl).Return(nil)

	// Start test
	target := NewTaskService(tdvm, trm, brm)

	_, err := target.Create(ctx, groupID, boardID, task)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
