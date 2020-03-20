package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
	mock_validation "github.com/16francs/gran/api/todo/mock/application/validation"
	mock_service "github.com/16francs/gran/api/todo/mock/domain/service"
)

var current = time.Now()

func TestBoardApplication_Index(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	u := &domain.User{}

	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	bs := []*domain.Board{b}

	// Defined mocks
	brvm := mock_validation.NewMockBoardRequestValidation(ctrl)

	bsm := mock_service.NewMockBoardService(ctrl)
	bsm.EXPECT().Index(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2").Return(bs, nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)
	usm.EXPECT().IsContainInGroupIDs(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2", u).Return(true)

	// Start test
	target := NewBoardApplication(brvm, bsm, usm)
	want := bs

	got, err := target.Index(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardApplication_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	req := &request.CreateBoard{
		Name:            "テストグループ",
		GroupID:         "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Closed:          true,
		Thumbnail:       "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
	}

	ves := make([]*domain.ValidationError, 0)

	u := &domain.User{}

	b := &domain.Board{
		Name:            "テストグループ",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
	}

	// Defined mocks
	brvm := mock_validation.NewMockBoardRequestValidation(ctrl)
	brvm.EXPECT().CreateBoard(req).Return(ves)

	bsm := mock_service.NewMockBoardService(ctrl)
	bsm.EXPECT().Create(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2", b).Return(nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)
	usm.EXPECT().IsContainInGroupIDs(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2", u).Return(true)

	// Start test
	target := NewBoardApplication(brvm, bsm, usm)

	err := target.Create(ctx, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
