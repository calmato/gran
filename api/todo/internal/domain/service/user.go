package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
)

// UserService - UserServiceインターフェース
type UserService interface {
	Authentication(ctx context.Context) (*domain.User, error)
	IsContainInGroupIDs(ctx context.Context, groupID string, u *domain.User) bool
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService - UserServiceの生成
func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
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

func (us *userService) IsContainInGroupIDs(ctx context.Context, groupID string, u *domain.User) bool {
	for _, v := range u.GroupIDs {
		if groupID == v {
			return true
		}
	}

	return false
}
