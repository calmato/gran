package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran/api/todo/internal/domain"
	mock_repository "github.com/calmato/gran/api/todo/mock/domain/repository"
	mock_uploader "github.com/calmato/gran/api/todo/mock/domain/uploader"
	mock_validation "github.com/calmato/gran/api/todo/mock/domain/validation"
	"github.com/golang/mock/gomock"
)

func TestBoardService_Index(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID  string
		Expected []*domain.Board
	}{
		"ok": {
			GroupID: "group-id",
			Expected: []*domain.Board{
				{
					ID:              "board-index-board-id",
					Name:            "テストグループ",
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

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Index(ctx, testCase.GroupID).Return(testCase.Expected, nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

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

func TestBoardService_Show(t *testing.T) {
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
				Name:            "テストグループ",
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
		bl := &domain.BoardList{
			ID:        "board-show-boardlist-id",
			Name:      "テストボードリスト",
			Color:     "red",
			BoardID:   testCase.BoardID,
			TaskIDs:   make([]string, 0),
			CreatedAt: current,
			UpdatedAt: current,
		}

		bls := []*domain.BoardList{bl}

		ts := make([]*domain.Task, 0)

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Show(ctx, testCase.GroupID, testCase.BoardID).Return(testCase.Expected, nil)
		brm.EXPECT().IndexBoardList(ctx, testCase.GroupID, testCase.BoardID).Return(bls, nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)
		trm.EXPECT().IndexByBoardID(ctx, testCase.BoardID).Return(ts, nil)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.Show(ctx, testCase.GroupID, testCase.BoardID)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardService_Create(t *testing.T) {
	testCases := map[string]struct {
		Board *domain.Board
	}{
		"ok": {
			Board: &domain.Board{
				Name:            "テストグループ",
				IsClosed:        true,
				ThumbnailURL:    "",
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

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)
		bdvm.EXPECT().Board(ctx, testCase.Board).Return(ves)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Create(ctx, testCase.Board).Return(nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.Create(ctx, testCase.Board)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.Board) {
				t.Fatalf("want %#v, but %#v", testCase.Board, got)
				return
			}
		})
	}
}

func TestBoardService_UploadThumbnail(t *testing.T) {
	testCases := map[string]struct {
		Data     []byte
		Expected string
	}{
		"ok": {
			Data:     []byte{},
			Expected: "http://localhost:8080",
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)
		fum.EXPECT().UploadBoardThumbnail(ctx, testCase.Data).Return(testCase.Expected, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.UploadThumbnail(ctx, testCase.Data)
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

func TestBoardService_Exists(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID  string
		BoardID  string
		Expected bool
	}{
		"ok": {
			GroupID:  "group-id",
			BoardID:  "board-id",
			Expected: true,
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		b := &domain.Board{
			ID:        testCase.BoardID,
			Name:      "テストボードリスト",
			CreatedAt: current,
			UpdatedAt: current,
		}

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Show(ctx, testCase.GroupID, testCase.BoardID).Return(b, nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got := target.Exists(ctx, testCase.GroupID, testCase.BoardID)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardService_ShowBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID     string
		BoardID     string
		BoardListID string
		Expected    *domain.BoardList
	}{
		"ok": {
			GroupID:     "group-id",
			BoardID:     "board-id",
			BoardListID: "board-list-id",
			Expected: &domain.BoardList{
				ID:    "board-list-id",
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

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().ShowBoardList(
			ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID,
		).Return(testCase.Expected, nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.ShowBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID)
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

func TestBoardService_CreateBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID   string
		BoardID   string
		BoardList *domain.BoardList
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			BoardList: &domain.BoardList{
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

		b := &domain.Board{
			ID: testCase.BoardID,
		}

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)
		bdvm.EXPECT().BoardList(ctx, testCase.BoardList).Return(ves)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Show(ctx, testCase.GroupID, testCase.BoardID).Return(b, nil)
		brm.EXPECT().Update(ctx, b).Return(nil)
		brm.EXPECT().CreateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardList).Return(nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.CreateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardList)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.BoardList) {
				t.Fatalf("want %#v, but %#v", testCase.BoardList, got)
				return
			}
		})
	}
}

func TestBoardService_UpdateBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID   string
		BoardID   string
		BoardList *domain.BoardList
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			BoardList: &domain.BoardList{
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

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)
		bdvm.EXPECT().BoardList(ctx, testCase.BoardList).Return(ves)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().UpdateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardList).Return(nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got, err := target.UpdateBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardList)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.BoardList) {
				t.Fatalf("want %#v, but %#v", testCase.BoardList, got)
				return
			}
		})
	}
}

func TestBoardService_ExistsBoardList(t *testing.T) {
	testCases := map[string]struct {
		GroupID     string
		BoardID     string
		BoardListID string
		Expected    bool
	}{
		"ok": {
			GroupID:     "group-id",
			BoardID:     "board-id",
			BoardListID: "board-list-id",
			Expected:    true,
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		bl := &domain.BoardList{
			ID:    testCase.BoardListID,
			Name:  "テストボードリスト",
			Color: "",
		}

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().ShowBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID).Return(bl, nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			got := target.ExistsBoardList(ctx, testCase.GroupID, testCase.BoardID, testCase.BoardListID)

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardService_UpdateKanban(t *testing.T) {
	testCases := map[string]struct {
		GroupID string
		BoardID string
		Board   *domain.Board
	}{
		"ok": {
			GroupID: "group-id",
			BoardID: "board-id",
			Board: &domain.Board{
				Name:    "テストボード",
				ListIDs: []string{"board-list-id"},
				Lists: map[string]*domain.BoardList{
					"board-list-id": {
						ID:   "board-list-id",
						Name: "ボードリスト",
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

		// Defined mocks
		bdvm := mock_validation.NewMockBoardDomainValidation(ctrl)
		bdvm.EXPECT().Board(ctx, testCase.Board).Return(ves)
		bdvm.EXPECT().BoardList(ctx, testCase.Board.Lists["board-list-id"]).Return(ves)

		brm := mock_repository.NewMockBoardRepository(ctrl)
		brm.EXPECT().Update(ctx, testCase.Board).Return(nil)
		brm.EXPECT().UpdateBoardList(
			ctx, testCase.GroupID, testCase.BoardID, testCase.Board.Lists["board-list-id"],
		).Return(nil)

		trm := mock_repository.NewMockTaskRepository(ctrl)

		fum := mock_uploader.NewMockFileUploader(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardService(bdvm, brm, trm, fum)

			err := target.UpdateKanban(ctx, testCase.GroupID, testCase.BoardID, testCase.Board)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
