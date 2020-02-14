package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/user/internal/domain"
)

var current = time.Now()

type userRepositoryMock struct{}

func (urm *userRepositoryMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
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

func TestUserDomainValidation_User(t *testing.T) {
	target := NewUserDomainValidation(&userRepositoryMock{})

	want := []*domain.ValidationError{}

	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got := target.User(ctx, u)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
