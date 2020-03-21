package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
	mock_repository "github.com/16francs/gran/api/todo/mock/domain/repository"
	mock_uploader "github.com/16francs/gran/api/todo/mock/domain/uploader"
	mock_validation "github.com/16francs/gran/api/todo/mock/domain/validation"
	"github.com/golang/mock/gomock"
)

func TestBoardService_Index(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-index-group-id"

	b := &domain.Board{
		ID:              "board-index-board-id",
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	bs := []*domain.Board{b}

	// Defined mocks
	brvm := mock_validation.NewMockBoardDomainValidation(ctrl)

	brm := mock_repository.NewMockBoardRepository(ctrl)
	brm.EXPECT().Index(ctx, groupID).Return(bs, nil)

	trm := mock_repository.NewMockTaskRepository(ctrl)

	fum := mock_uploader.NewMockFileUploader(ctrl)

	// Start test
	target := NewBoardService(brvm, brm, trm, fum)

	want := bs

	got, err := target.Index(ctx, groupID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardService_Show(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-show-group-id"
	boardID := "board-show-board-id"

	b := &domain.Board{
		ID:              boardID,
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         groupID,
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	bl := &domain.BoardList{
		ID:        "board-show-boardlist-id",
		Name:      "テストボードリスト",
		Color:     "red",
		BoardID:   boardID,
		TaskIDs:   make([]string, 0),
		CreatedAt: current,
		UpdatedAt: current,
	}

	bls := []*domain.BoardList{bl}

	// Defined mocks
	brvm := mock_validation.NewMockBoardDomainValidation(ctrl)

	brm := mock_repository.NewMockBoardRepository(ctrl)
	brm.EXPECT().Show(ctx, groupID, boardID).Return(b, nil)
	brm.EXPECT().IndexBoardList(ctx, groupID, boardID).Return(bls, nil)

	trm := mock_repository.NewMockTaskRepository(ctrl)
	trm.EXPECT().IndexByBoardListID(ctx, bl.ID).Return([]*domain.Task{}, nil)

	fum := mock_uploader.NewMockFileUploader(ctrl)

	// Start test
	target := NewBoardService(brvm, brm, trm, fum)

	want := b
	want.Lists = bls
	want.Lists[0].Tasks = []*domain.Task{}

	got, err := target.Show(ctx, groupID, boardID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardService_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	ves := make([]*domain.ValidationError, 0)
	groupID := "board-create-group-id"

	b := &domain.Board{
		ID:              "board-create-board-id",
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	// Defined mocks
	brvm := mock_validation.NewMockBoardDomainValidation(ctrl)
	brvm.EXPECT().Board(ctx, b).Return(ves)

	brm := mock_repository.NewMockBoardRepository(ctrl)
	brm.EXPECT().Create(ctx, b).Return(nil)

	trm := mock_repository.NewMockTaskRepository(ctrl)

	fum := mock_uploader.NewMockFileUploader(ctrl)

	// Start test
	target := NewBoardService(brvm, brm, trm, fum)

	_, err := target.Create(ctx, groupID, b)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestBoardService_UploadThumbnail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	thumbnailURL := "http://localhost:8080"

	// Defined mocks
	brvm := mock_validation.NewMockBoardDomainValidation(ctrl)

	brm := mock_repository.NewMockBoardRepository(ctrl)

	trm := mock_repository.NewMockTaskRepository(ctrl)

	fum := mock_uploader.NewMockFileUploader(ctrl)
	fum.EXPECT().UploadBoardThumbnail(ctx, []byte{}).Return(thumbnailURL, nil)

	// Start test
	target := NewBoardService(brvm, brm, trm, fum)

	want := thumbnailURL

	got, err := target.UploadThumbnail(ctx, []byte{})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
