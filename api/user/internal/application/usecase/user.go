package usecase

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
)

// UserUsecase - UserUsecaseインターフェース
type UserUsecase interface {
	Create(ctx context.Context, req request.CreateUser) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase - UserUsecaseの生成
func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) Create(ctx context.Context, req request.CreateUser) error {
	// TODO: validation check

	current := time.Now()
	u := &domain.User{
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: current,
		UpdatedAt: current,
	}

	if err := uu.userRepository.Create(ctx, u); err != nil {
		return xerrors.Errorf("Failed to UserUsecase/Create: %w", err)
	}

	return nil
}
