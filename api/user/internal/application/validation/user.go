package validation

import (
	"github.com/16francs/gran/api/user/internal/application/request"
)

// UserValidation - ユーザー関連のバリデーション
type UserValidation interface {
	CreateUser(cu request.CreateUser) error
}

type userValidation struct {
	validator CustomValidator
}

// NewUserValidation - UserValidationの生成
func NewUserValidation() UserValidation {
	v := NewCustomValidator()

	return &userValidation{v}
}

func (uv *userValidation) CreateUser(cu request.CreateUser) error {
	return uv.validator.Run(cu)
}
