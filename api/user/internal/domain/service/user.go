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
	Create(ctx context.Context, u *domain.User) error
	CreateGroup(ctx context.Context, u *domain.User, g *domain.Group) error
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

func (us *userService) Create(ctx context.Context, u *domain.User) error {
	if err := us.userDomainValidation.User(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/DomainValidation: %w", err)
		return domain.InvalidDomainValidation.New(err)
	}

	if err := us.userRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.Unknown.New(err)
	}

	return nil
}

func (us *userService) CreateGroup(ctx context.Context, u *domain.User, g *domain.Group) error {
	if err := us.userDomainValidation.Group(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/DomainValidation: %w", err)
		return domain.InvalidDomainValidation.New(err)
	}

	u.Groups = append(u.Groups, *g)

	if err := us.userRepository.CreateGroup(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.Unknown.New(err)
	}

	return nil
}
