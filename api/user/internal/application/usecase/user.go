package usecase

import (
	"context"
	"time"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
)

// UserUsecase - UserUsecaseインターフェース
type UserUsecase interface {
	Create(ctx context.Context, req request.CreateUser) error
}

type userUsecase struct {
	userValidation validation.UserValidation
	userRepository repository.UserRepository
}

// NewUserUsecase - UserUsecaseの生成
func NewUserUsecase(uv validation.UserValidation, ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userValidation: uv,
		userRepository: ur,
	}
}

func (uu *userUsecase) Create(ctx context.Context, req request.CreateUser) error {
	if err := uu.userValidation.CreateUser(req); err != nil {
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

	if err := uu.userRepository.Create(ctx, u); err != nil {
		return err
	}

	return nil
}
