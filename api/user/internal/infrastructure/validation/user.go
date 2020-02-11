package validation

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	dv "github.com/16francs/gran/api/user/internal/domain/validation"
)

type userDomainValidation struct {
	userRepository repository.UserRepository
}

// NewUserDomainValidation - UserDomainValidationの生成
func NewUserDomainValidation(ur repository.UserRepository) dv.UserDomainValidation {
	return &userDomainValidation{
		userRepository: ur,
	}
}

func (udv *userDomainValidation) User(ctx context.Context, u *domain.User) []*domain.ValidationError {
	validationErrors := make([]*domain.ValidationError, 0)

	if err := uniqueCheckEmail(ctx, udv.userRepository, u.Email); err != nil {
		validationError := &domain.ValidationError{
			Field:   "メールアドレス",
			Message: dv.CustomUniqueMessage,
		}

		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

func uniqueCheckEmail(ctx context.Context, ur repository.UserRepository, email string) error {
	uid, _ := ur.GetUIDByEmail(ctx, email)
	if uid != "" {
		return xerrors.New("This email is already exists.")
	}

	return nil
}
