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
	current := time.Now()

	testCases := map[string]struct {
		Expected *domain.User
	}{
		"ok": {
			Expected: &domain.User{
				ID:           "user-id",
				Email:        "hoge@hoge.com",
				Password:     "12345678",
				Name:         "",
				ThumbnailURL: "",
				GroupIDs:     make([]string, 0),
				CreatedAt:    current,
				UpdatedAt:    current,
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined mocks
		udvm := mock_validation.NewMockUserDomainValidation(ctrl)

		urm := mock_repository.NewMockUserRepository(ctrl)
		urm.EXPECT().Authentication(ctx).Return(testCase.Expected, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserService(udvm, urm)

			got, err := target.Authentication(ctx)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	testCases := map[string]struct {
		User *domain.User
	}{
		"ok": {
			User: &domain.User{
				Email:        "hoge@hoge.com",
				Password:     "12345678",
				Name:         "",
				ThumbnailURL: "",
				GroupIDs:     make([]string, 0),
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		// Defined mocks
		udvm := mock_validation.NewMockUserDomainValidation(ctrl)
		udvm.EXPECT().User(ctx, testCase.User).Return(ves)

		urm := mock_repository.NewMockUserRepository(ctrl)
		urm.EXPECT().Create(ctx, testCase.User).Return(nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserService(udvm, urm)

			got, err := target.Create(ctx, testCase.User)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.User) {
				t.Fatalf("want %#v, but %#v", testCase.User, got)
				return
			}
		})
	}
}
