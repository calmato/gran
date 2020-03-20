package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/todo/internal/domain"
	mock_repository "github.com/16francs/gran/api/todo/mock/domain/repository"
)

func TestUserService_Authentication(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()

	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	brm := mock_repository.NewMockUserRepository(ctrl)
	brm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewUserService(brm)

	want := u

	got, err := target.Authentication(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestUserService_IsContainInGroupIDs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "JUA1ouY12ickxIupMVdVl3ieM7s2"

	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     []string{groupID},
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	brm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewUserService(brm)

	got := target.IsContainInGroupIDs(ctx, groupID, u)
	if !got {
		t.Fatalf("error: %v", got)
	}
}
