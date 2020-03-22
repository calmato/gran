package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/user/internal/domain"
	mock_repository "github.com/16francs/gran/api/user/mock/domain/repository"
)

func TestUserDomainValidation_User(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()

	u := &domain.User{
		ID:           "user-id",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	// Defined mocks
	urm := mock_repository.NewMockUserRepository(ctrl)
	urm.EXPECT().GetUIDByEmail(ctx, u.Email).Return("", nil)

	// Start test
	target := NewUserDomainValidation(urm)

	want := []*domain.ValidationError{}

	got := target.User(ctx, u)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
