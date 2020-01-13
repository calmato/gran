package validation

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
)

// UserDomainValidation - UserDomainRepositoryインターフェース
type UserDomainValidation interface {
	User(ctx context.Context, u *domain.User) error
	Group(ctx context.Context, g *domain.Group) error
}
