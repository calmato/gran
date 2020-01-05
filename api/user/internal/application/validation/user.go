package validation

import "github.com/16francs/gran/api/user/internal/application/request"

// UserRequestValidation - ユーザー関連のバリデーション
type UserRequestValidation interface {
	CreateUser(cu request.CreateUser) error
}

type userRequestValidation struct {
	validator RequestValidator
}

// NewUserRequestValidation - UserRequestValidationの生成
func NewUserRequestValidation() UserRequestValidation {
	v := NewRequestValidator()

	return &userRequestValidation{v}
}

func (urv *userRequestValidation) CreateUser(cu request.CreateUser) error {
	return urv.validator.Run(cu)
}
