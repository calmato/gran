package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/user/internal/domain"
	mock_repository "github.com/16francs/gran/api/user/mock/domain/repository"
	mock_validation "github.com/16francs/gran/api/user/mock/domain/validation"
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
	udvm := mock_validation.NewMockUserDomainValidation(ctrl)

	urm := mock_repository.NewMockUserRepository(ctrl)
	urm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewUserService(udvm, urm)

	want := u

	got, err := target.Authentication(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestUserService_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)
	current := time.Now()

	u := &domain.User{
		ID:           "user-create-user-id",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	udvm := mock_validation.NewMockUserDomainValidation(ctrl)
	udvm.EXPECT().User(ctx, u).Return(ves)

	urm := mock_repository.NewMockUserRepository(ctrl)
	urm.EXPECT().Create(ctx, u).Return(nil)

	// Start test
	target := NewUserService(udvm, urm)

	_, err := target.Create(ctx, u)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
