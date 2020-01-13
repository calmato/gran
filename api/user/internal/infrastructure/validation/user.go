package validation

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	dv "github.com/16francs/gran/api/user/internal/domain/validation"
)

type userDomainValidation struct {
	validator      DomainValidator
	userRepository repository.UserRepository
}

// NewUserDomainValidation - UserDomainValidationの生成
func NewUserDomainValidation(ur repository.UserRepository) dv.UserDomainValidation {
	v := NewDomainValidator()

	return &userDomainValidation{
		validator:      v,
		userRepository: ur,
	}
}

func (udv *userDomainValidation) User(ctx context.Context, u *domain.User) error {
	err := udv.validator.Run(u)
	if err != nil {
		return err
	}

	err = uniqueCheckEmail(ctx, udv.userRepository, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (udv *userDomainValidation) Group(ctx context.Context, g *domain.Group) error {
	return udv.validator.Run(g)
}

func uniqueCheckEmail(ctx context.Context, ur repository.UserRepository, email string) error {
	uid, _ := ur.GetUIDByEmail(ctx, email)
	if uid != "" {
		return xerrors.New("Email is not unique.")
	}

	return nil
}
