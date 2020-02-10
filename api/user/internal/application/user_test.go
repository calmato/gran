package application

import (
	"context"
	"testing"
	"time"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
)

var current = time.Now()

type userRequestValidationMock struct{}

func (urvm *userRequestValidationMock) CreateUser(req *request.CreateUser) []*domain.ValidationError {
	return nil
}

type userServiceMock struct{}

func (usm *userServiceMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	return u, nil
}

func (usm *userServiceMock) Create(ctx context.Context, u *domain.User) error {
	return nil
}

func TestUserApplication_Create(t *testing.T) {
	target := NewUserApplication(&userRequestValidationMock{}, &userServiceMock{})

	u := &request.CreateUser{
		Email:                "hoge@hoge.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Create(ctx, u)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
