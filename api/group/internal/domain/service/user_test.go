package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/group/internal/domain"
	mock_repository "github.com/16francs/gran/api/group/mock/domain/repository"
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
		urm := mock_repository.NewMockUserRepository(ctrl)
		urm.EXPECT().Authentication(ctx).Return(testCase.Expected, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserService(urm)

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
