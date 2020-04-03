package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
	mock_repository "github.com/16francs/gran/api/todo/mock/domain/repository"
	mock_validation "github.com/16francs/gran/api/todo/mock/domain/validation"
	"github.com/golang/mock/gomock"
)

func TestTaskService_Create(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID     string
		BoardID     string
		BoardListID string
		Task        *domain.Task
	}{
		"ok": {
			GroupID:     "group-id",
			BoardID:     "board-id",
			BoardListID: "board-list-id",
			Task: &domain.Task{
				Name:            "タスク",
				Description:     "説明",
				Labels:          []string{},
				AttachmentURLs:  []string{},
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

		bl := &domain.BoardList{
			ID:      testCase.BoardListID,
			Name:    "ボードリスト",
			Color:   "",
			TaskIDs: []string{},
		}

		// Defined mocks
		tdvm := mock_validation.NewMockTaskDomainValidation(ctrl)
		tdvm.EXPECT().Task(ctx, testCase.Task).Return(ves)

		trm := mock_repository.NewMockTaskRepository(ctrl)
		trm.EXPECT().Create(ctx, testCase.Task).Return(nil)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().ShowBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID).Return(bl, nil)
		brm.EXPECT().UpdateBoardList(ctx, testCase.GroupID, testCase.BoardID, bl).Return(nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskService(tdvm, trm, brm)

			got, err := target.Create(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID, testCase.Task)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.Task) {
				t.Fatalf("want %#v, but %#v", testCase.Task, got)
				return
			}
		})
	}
}
