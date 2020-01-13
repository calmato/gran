package application

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/service"
)

// UserApplication - UserApplicationインターフェース
type UserApplication interface {
	Create(ctx context.Context, req *request.CreateUser) error
	CreateGroup(ctx context.Context, req *request.CreateGroup) error
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
	if err := ua.userRequestValidation.CreateUser(req); err != nil {
		err = xerrors.Errorf("Failed to Application/RequestValidation: %w", err)
		return domain.InvalidRequestValidation.New(err)
	}

	current := time.Now()
	u := &domain.User{
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: current,
		UpdatedAt: current,
	}

	if err := ua.userService.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

func (ua *userApplication) CreateGroup(ctx context.Context, req *request.CreateGroup) error {
	// TODO: 認証処理
	u := &domain.User{}

	if err := ua.userRequestValidation.CreateGroup(req); err != nil {
		err = xerrors.Errorf("Failed to Application/RequestValidation: %w", err)
		return domain.InvalidRequestValidation.New(err)
	}

	current := time.Now()
	g := &domain.Group{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	if err := ua.userService.CreateGroup(ctx, u, g); err != nil {
		return err
	}

	return nil
}
