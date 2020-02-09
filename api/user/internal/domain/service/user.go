package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
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
	if ves := us.userDomainValidation.User(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")

		if containUniqueError(ves) {
			return domain.AlreadyExists.New(err, ves...)
		}

		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := us.userRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func containUniqueError(ves []*domain.ValidationError) bool {
	for _, v := range ves {
		if v.Message == validation.CustomUniqueMessage {
			return true
		}
	}

	return false
}
