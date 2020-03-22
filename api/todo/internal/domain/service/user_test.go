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
		ID:           "user-authentication-user-id",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	urm := mock_repository.NewMockUserRepository(ctrl)
	urm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewUserService(urm)

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
	groupID := "user-iscontainingroupids-group-id"

	u := &domain.User{
		ID:           "user-iscontaingroupids-user-id",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     []string{groupID},
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewUserService(urm)

	got := target.IsContainInGroupIDs(ctx, groupID, u)
	if !got {
		t.Fatalf("error: %v", got)
	}
}
