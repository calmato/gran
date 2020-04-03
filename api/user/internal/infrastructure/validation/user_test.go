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
	current := time.Now()

	testCases := map[string]struct {
		User     *domain.User
		Expected []*domain.ValidationError
	}{
		"ok": {
			User: &domain.User{
				ID:           "user-id",
				Email:        "hoge@hoge.com",
				Password:     "12345678",
				Name:         "",
				ThumbnailURL: "",
				GroupIDs:     make([]string, 0),
				CreatedAt:    current,
				UpdatedAt:    current,
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined mocks
		urm := mock_repository.NewMockUserRepository(ctrl)
		urm.EXPECT().GetUIDByEmail(ctx, testCase.User.Email).Return("", nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserDomainValidation(urm)

			got := target.User(ctx, testCase.User)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
