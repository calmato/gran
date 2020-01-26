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
	Authentication(ctx context.Context) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
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

func (us *userService) Authentication(ctx context.Context) (*domain.User, error) {
	u, err := us.userRepository.Authentication(ctx)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return u, nil
}

func (us *userService) Create(ctx context.Context, u *domain.User) error {
	if err := us.userDomainValidation.User(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/DomainValidation: %w", err)
		return domain.InvalidDomainValidation.New(err)
	}

	if err := us.userRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}
