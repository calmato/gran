package service

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/internal/domain/validation"
)

// UserService - UserServiceインターフェース
type UserService interface {
	Authentication(ctx context.Context) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) (*domain.User, error)
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

func (us *userService) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	if ves := us.userDomainValidation.User(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")

		if isContainCustomUniqueError(ves) {
			return nil, domain.AlreadyExists.New(err, ves...)
		}

		return nil, domain.Unknown.New(err, ves...)
	}

	current := time.Now()

	u.CreatedAt = current
	u.UpdatedAt = current

	if err := us.userRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return u, nil
}

func (us *userService) Update(ctx context.Context, u *domain.User) (*domain.User, error) {
	if ves := us.userDomainValidation.User(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")

		if isContainCustomUniqueError(ves) {
			return nil, domain.AlreadyExists.New(err, ves...)
		}

		return nil, domain.Unknown.New(err, ves...)
	}
	current := time.Now()

	u.UpdatedAt = current

	if err := us.userRepository.Update(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return u, nil
}

func isContainCustomUniqueError(ves []*domain.ValidationError) bool {
	for _, v := range ves {
		if v.Message == validation.CustomUniqueMessage {
			return true
		}
	}

	return false
}
