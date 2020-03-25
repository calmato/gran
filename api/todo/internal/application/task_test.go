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
	current := time.Now()

	testCases := map[string]struct {
		GroupID string
		BoardID string
		Request *request.CreateTask
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			Request: &request.CreateTask{
				Name:            "タスク",
				Description:     "説明",
				BoardListID:     "board-list-id",
				Labels:          []string{},
				Attachments:     []string{},
				AssignedUserIDs: []string{},
				DeadlinedAt:     current,
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{}

		task := &domain.Task{
			Name:            testCase.Request.Name,
			Description:     testCase.Request.Description,
			Labels:          testCase.Request.Labels,
			AssignedUserIDs: testCase.Request.AssignedUserIDs,
			DeadlinedAt:     testCase.Request.DeadlinedAt,
			AttachmentURLs:  make([]string, 0),
			BoardID:         testCase.BoardID,
			BoardListID:     testCase.Request.BoardListID,
		}

		// Defined mocks
		trvm := mock_validation.NewMockTaskRequestValidation(ctrl)
		trvm.EXPECT().CreateTask(testCase.Request).Return(ves)

		tsm := mock_service.NewMockTaskService(ctrl)
		tsm.EXPECT().Create(ctx, testCase.GroupID, testCase.BoardID, task).Return(task, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskApplication(trvm, tsm, usm)

			err := target.Create(ctx, testCase.GroupID, testCase.BoardID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
