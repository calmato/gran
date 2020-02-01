package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/user/internal/domain"
)

var current = time.Now()

type userDomainValidationMock struct{}

func (udvm *userDomainValidationMock) User(ctx context.Context, u *domain.User) []*domain.ValidationError {
	return nil
}

type userRepositoryMock struct{}

func (urm *userRepositoryMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	return u, nil
}

func (urm *userRepositoryMock) Create(ctx context.Context, u *domain.User) error {
	return nil
}

func (urm *userRepositoryMock) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	return "", nil
}

func TestUserService_Authentication(t *testing.T) {
	target := NewUserService(&userDomainValidationMock{}, &userRepositoryMock{})

	want := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	got, err := target.Authentication(nil)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestUserService_Create(t *testing.T) {
	target := NewUserService(&userDomainValidationMock{}, &userRepositoryMock{})

	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	err := target.Create(nil, u)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
