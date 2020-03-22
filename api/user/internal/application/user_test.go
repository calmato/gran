package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
	mock_validation "github.com/16francs/gran/api/user/mock/application/validation"
	mock_service "github.com/16francs/gran/api/user/mock/domain/service"
)

func TestUserApplication_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	req := &request.CreateUser{
		Email:                "hoge@hoge.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	u := &domain.User{
		Email:    "hoge@hoge.com",
		Password: "12345678",
	}

	// Defined mocks
	urvm := mock_validation.NewMockUserRequestValidation(ctrl)
	urvm.EXPECT().CreateUser(req).Return(ves)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Create(ctx, u).Return(nil)

	target := NewUserApplication(urvm, usm)

	err := target.Create(ctx, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
