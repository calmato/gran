package validation

import "github.com/16francs/gran/api/user/internal/domain"

// UserDomainValidation - UserDomainRepositoryインターフェース
type UserDomainValidation interface {
	User(u *domain.User) error
}
