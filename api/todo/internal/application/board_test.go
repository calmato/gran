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

func TestBoardApplication_Index(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID  string
		Expected []*domain.Board
	}{
		"ok": {
			GroupID: "board-id",
			Expected: []*domain.Board{
				{
					ID:              "board-id",
					Name:            "テストボード",
					IsClosed:        true,
					ThumbnailURL:    "",
					BackgroundColor: "",
					Labels:          make([]string, 0),
					GroupID:         "",
					CreatedAt:       current,
					UpdatedAt:       current,
				},
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
			GroupIDs: []string{testCase.GroupID},
		}

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().Index(ctx, testCase.GroupID).Return(testCase.Expected, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			got, err := target.Index(ctx, testCase.GroupID)
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

func TestBoardApplication_Show(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID  string
		BoardID  string
		Expected *domain.Board
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			Expected: &domain.Board{
				ID:              "board-id",
				Name:            "テストボード",
				IsClosed:        true,
				ThumbnailURL:    "",
				BackgroundColor: "",
				Labels:          make([]string, 0),
				GroupID:         "group-id",
				CreatedAt:       current,
				UpdatedAt:       current,
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
			ID:       "user-id",
			GroupIDs: []string{testCase.GroupID},
		}

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().Show(ctx, testCase.GroupID, testCase.BoardID).Return(testCase.Expected, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			got, err := target.Show(ctx, testCase.GroupID, testCase.BoardID)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
			}
		})
	}
}

func TestBoardApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		GroupID string
		Request *request.CreateBoard
	}{
		"ok": {
			GroupID: "group-id",
			Request: &request.CreateBoard{
				Name:            "テストグループ",
				IsClosed:        true,
				Thumbnail:       "",
				BackgroundColor: "",
				Labels:          make([]string, 0),
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

		u := &domain.User{
			GroupIDs: []string{testCase.GroupID},
		}

		b := &domain.Board{
			Name:            testCase.Request.Name,
			IsClosed:        testCase.Request.IsClosed,
			ThumbnailURL:    "",
			BackgroundColor: testCase.Request.BackgroundColor,
			Labels:          testCase.Request.Labels,
			GroupID:         testCase.GroupID,
		}

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)
		brvm.EXPECT().CreateBoard(testCase.Request).Return(ves)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().Create(ctx, b).Return(b, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			err := target.Create(ctx, testCase.GroupID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestBoardApplication_CreateBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID string
		BoardID string
		Request *request.CreateBoardList
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			Request: &request.CreateBoardList{
				Name:  "テストボードリスト",
				Color: "",
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

		u := &domain.User{
			GroupIDs: []string{testCase.GroupID},
		}

		bl := &domain.BoardList{
			Name:    testCase.Request.Name,
			Color:   testCase.Request.Color,
			BoardID: testCase.BoardID,
			TaskIDs: make([]string, 0),
		}

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)
		brvm.EXPECT().CreateBoardList(testCase.Request).Return(ves)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().Exists(ctx, testCase.GroupID, testCase.BoardID).Return(true)
		bsm.EXPECT().CreateBoardList(ctx, testCase.GroupID, testCase.BoardID, bl).Return(bl, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			_, err := target.CreateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestBoardApplication_UpdateBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID     string
		BoardID     string
		BoardListID string
		Request     *request.UpdateBoardList
	}{
		"ok": {
			GroupID:     "group-id",
			BoardID:     "board-id",
			BoardListID: "board-list-id",
			Request: &request.UpdateBoardList{
				Name:  "テストボードリスト",
				Color: "",
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

		u := &domain.User{
			GroupIDs: []string{testCase.GroupID},
		}

		bl := &domain.BoardList{
			ID:      testCase.BoardListID,
			Name:    testCase.Request.Name,
			Color:   testCase.Request.Color,
			BoardID: testCase.BoardID,
			TaskIDs: make([]string, 0),
		}

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)
		brvm.EXPECT().UpdateBoardList(testCase.Request).Return(ves)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().ShowBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID).Return(bl, nil)
		bsm.EXPECT().UpdateBoardList(ctx, testCase.GroupID, testCase.BoardID, bl).Return(bl, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			err := target.UpdateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestBoardApplication_UpdateKanban(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID string
		BoardID string
		Request *request.UpdateKanban
		Board   *domain.Board
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			Request: &request.UpdateKanban{
				Lists: []*request.ListInUpdateKanban{
					{
						ID:    "board-list-id-01",
						Name:  "テストボードリスト01",
						Color: "red",
						Tasks: []*request.TaskInUpdateKanban{
							{
								ID:              "task-id",
								Name:            "テストタスク",
								Labels:          make([]string, 0),
								AssignedUserIDs: make([]string, 0),
								DeadlinedAt:     current,
							},
						},
					},
					{
						ID:    "board-list-id-02",
						Name:  "テストボードリスト02",
						Color: "red",
						Tasks: make([]*request.TaskInUpdateKanban, 0),
					},
				},
			},
			Board: &domain.Board{
				ID: "board-id",
				ListIDs: []string{
					"board-list-id-01",
					"board-list-id-02",
				},
				Lists: map[string]*domain.BoardList{
					"board-list-id-01": {
						ID:      "board-list-id-01",
						Name:    "テストボードリスト01",
						Color:   "red",
						TaskIDs: make([]string, 0),
						Tasks:   map[string]*domain.Task{},
					},
					"board-list-id-02": {
						ID:      "board-list-id-02",
						Name:    "テストボードリスト02",
						Color:   "red",
						TaskIDs: []string{"task-id"},
						Tasks: map[string]*domain.Task{
							"task-id": {
								ID:              "task-id",
								Name:            "テストタスク",
								Labels:          make([]string, 0),
								AssignedUserIDs: make([]string, 0),
								DeadlinedAt:     current,
							},
						},
					},
				},
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

		u := &domain.User{
			GroupIDs: []string{testCase.GroupID},
		}

		listIDs := make([]string, len(testCase.Request.Lists))
		for i, bl := range testCase.Request.Lists {
			listIDs[i] = bl.ID

			taskIDs := make([]string, len(bl.Tasks))
			for j, t := range bl.Tasks {
				taskIDs[j] = t.ID
			}

			testCase.Board.Lists[bl.ID].TaskIDs = taskIDs
		}

		testCase.Board.ListIDs = listIDs

		// Defined mocks
		brvm := mock_validation.NewMockBoardRequestValidation(ctrl)
		brvm.EXPECT().UpdateKanban(testCase.Request).Return(ves)

		bsm := mock_service.NewMockBoardService(ctrl)
		bsm.EXPECT().Show(ctx, testCase.GroupID, testCase.BoardID).Return(testCase.Board, nil)
		bsm.EXPECT().UpdateKanban(ctx, testCase.GroupID, testCase.BoardID, testCase.Board).Return(nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().IsContainInGroupIDs(ctx, testCase.GroupID, u).Return(true)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardApplication(brvm, bsm, usm)

			err := target.UpdateKanban(ctx, testCase.GroupID, testCase.BoardID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
