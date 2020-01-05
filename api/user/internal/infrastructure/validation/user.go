package validation

import (
	"github.com/16francs/gran/api/user/internal/domain"
	dv "github.com/16francs/gran/api/user/internal/domain/validation"
)

type userDomainValidation struct {
	validator DomainValidator
}

// NewUserDomainValidation - UserDomainValidationの生成
func NewUserDomainValidation() dv.UserDomainValidation {
	v := NewDomainValidator()

	return &userDomainValidation{v}
}

func (udv *userDomainValidation) User(u *domain.User) error {
	return udv.validator.Run(u)
}
