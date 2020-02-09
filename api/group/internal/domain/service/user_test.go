package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
)

var userCurrent = time.Now()

type userRepositoryMock struct{}

func (urm *userRepositoryMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    userCurrent,
		UpdatedAt:    userCurrent,
	}

	return u, nil
}

func TestUserService_Authentication(t *testing.T) {
	target := NewUserService(&userRepositoryMock{})

	want := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    userCurrent,
		UpdatedAt:    userCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.Authentication(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
