package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain/repository"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/validation"
)

// UserService - UserServiceインターフェース
type UserService interface {
	Create(ctx context.Context, u domain.User) error
}

type userService struct {
	userDomainValidation validation.UserDomainValidation
	userRepository       repository.UserRepository
}

// NewUserService - UserServiceの生成
func NewUserService(udv validation.UserDomainValidation, ur repository.UserRepository) UserService {
	return &userService{
		userDomainValidation: udv,
		userRepository:       ur,
	}
}

func (us *userService) Create(ctx context.Context, u domain.User) error {
	if err := us.userDomainValidation.User(&u); err != nil {
		return xerrors.Errorf("Failed to UserService/Create: %w", err)
	}

	if err := us.userRepository.Create(ctx, &u); err != nil {
		return xerrors.Errorf("Failed to UserService/Create: %w", err)
	}

	return nil
}
