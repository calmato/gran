package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/calmato/gran/api/todo/internal/application/request"
	"github.com/calmato/gran/api/todo/internal/domain"
	mock_validation "github.com/calmato/gran/api/todo/mock/application/validation"
	mock_service "github.com/calmato/gran/api/todo/mock/domain/service"
)

func TestTaskApplication_Show(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		TaskID   string
		Expected *domain.Task
	}{
		"ok": {
			TaskID: "task-id",
			Expected: &domain.Task{
				ID:              "task-id",
				Name:            "タスク",
				Description:     "説明",
				BoardListID:     "board-list-id",
				Labels:          []string{},
				AttachmentURLs:  []string{},
				AssignedUserIDs: []string{},
				DeadlinedAt:     current,
				GroupID:         "group-id",
				BoardID:         "board-id",
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		u := &domain.User{
			GroupIDs: []string{testCase.Expected.GroupID},
		}

		// Defined mocks
		trvm := mock_validation.NewMockTaskRequestValidation(ctrl)

		tsm := mock_service.NewMockTaskService(ctrl)
		tsm.EXPECT().Show(ctx, testCase.TaskID).Return(testCase.Expected, nil)

		bsm := mock_service.NewMockBoardService(ctrl)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.Expected.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskApplication(trvm, tsm, bsm, usm)

			got, err := target.Show(ctx, testCase.TaskID)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

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
			GroupID:         testCase.GroupID,
			BoardID:         testCase.BoardID,
		}

		// Defined mocks
		trvm := mock_validation.NewMockTaskRequestValidation(ctrl)
		trvm.EXPECT().CreateTask(testCase.Request).Return(ves)

		tsm := mock_service.NewMockTaskService(ctrl)
		tsm.EXPECT().Create(ctx, testCase.GroupID, testCase.BoardID, testCase.Request.BoardListID, task).Return(task, nil)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().ExistsBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.Request.BoardListID).Return(true)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskApplication(trvm, tsm, bsm, usm)

			_, err := target.Create(ctx, testCase.GroupID, testCase.BoardID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
