package application

import (
	"context"
	"time"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
)

// UserApplication - UserApplicationインターフェース
type UserApplication interface {
	Create(ctx context.Context, req request.CreateUser) error
}

type userApplication struct {
	userRequestValidation validation.UserRequestValidation
	userRepository        repository.UserRepository
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(urv validation.UserRequestValidation, ur repository.UserRepository) UserApplication {
	return &userApplication{
		userRequestValidation: urv,
		userRepository:        ur,
	}
}

func (ua *userApplication) Create(ctx context.Context, req request.CreateUser) error {
	if err := ua.userRequestValidation.CreateUser(req); err != nil {
		return err // TODO: エラーメッセージをレスポンスに
	}

	current := time.Now()
	u := &domain.User{
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: current,
		UpdatedAt: current,
	}

	// TODO: ユニークチェック

	if err := ua.userRepository.Create(ctx, u); err != nil {
		return err
	}

	return nil
}
