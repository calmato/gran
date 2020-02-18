package application

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/service"
)

// UserApplication - UserApplicationインターフェース
type UserApplication interface {
	Create(ctx context.Context, req *request.CreateUser) error
}

type userApplication struct {
	userRequestValidation validation.UserRequestValidation
	userService           service.UserService
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(urv validation.UserRequestValidation, us service.UserService) UserApplication {
	return &userApplication{
		userRequestValidation: urv,
		userService:           us,
	}
}

func (ua *userApplication) Create(ctx context.Context, req *request.CreateUser) error {
	if ves := ua.userRequestValidation.CreateUser(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	u := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := ua.userService.Create(ctx, u); err != nil {
		return err
	}

	return nil
}
