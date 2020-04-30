package validation

import (
	"context"

	"github.com/calmato/gran/api/user/internal/domain"
)

// UserDomainValidation - UserDomainRepositoryインターフェース
type UserDomainValidation interface {
	User(ctx context.Context, u *domain.User) []*domain.ValidationError
}
