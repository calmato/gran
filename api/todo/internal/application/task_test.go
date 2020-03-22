package application

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
	mock_validation "github.com/16francs/gran/api/todo/mock/application/validation"
	mock_service "github.com/16francs/gran/api/todo/mock/domain/service"
)

func TestTaskApplication_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "task-create-group-id"
	boardID := "task-create-board-id"
	ves := make([]*domain.ValidationError, 0)

	req := &request.CreateTask{
		Name:            "タスク",
		Description:     "説明",
		BoardListID:     "task-create-boardlist-id",
		Labels:          []string{},
		Attachments:     []string{},
		AssignedUserIDs: []string{},
		DeadlinedAt:     current,
	}

	u := &domain.User{}

	task := &domain.Task{
		Name:            "タスク",
		Description:     "説明",
		BoardListID:     "task-create-boardlist-id",
		Labels:          []string{},
		AttachmentURLs:  []string{},
		AssignedUserIDs: []string{},
		DeadlinedAt:     current,
	}

	// Defined mocks
	trvm := mock_validation.NewMockTaskRequestValidation(ctrl)
	trvm.EXPECT().CreateTask(req).Return(ves)

	tsm := mock_service.NewMockTaskService(ctrl)
	tsm.EXPECT().Create(ctx, boardID, task).Return(task, nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)
	usm.EXPECT().IsContainInGroupIDs(ctx, groupID, u).Return(true)

	// Start test
	target := NewTaskApplication(trvm, tsm, usm)

	err := target.Create(ctx, groupID, boardID, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
