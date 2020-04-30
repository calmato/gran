package validation

import (
	"github.com/calmato/gran/api/user/internal/application/request"
	"github.com/calmato/gran/api/user/internal/domain"
)

// UserRequestValidation - ユーザー関連のバリデーション
type UserRequestValidation interface {
	CreateUser(req *request.CreateUser) []*domain.ValidationError
	UpdateProfile(req *request.UpdateProfile) []*domain.ValidationError
}

type userRequestValidation struct {
	validator RequestValidator
}

// NewUserRequestValidation - UserRequestValidationの生成
func NewUserRequestValidation() UserRequestValidation {
	rv := NewRequestValidator()

	return &userRequestValidation{
		validator: rv,
	}
}

func (urv *userRequestValidation) CreateUser(req *request.CreateUser) []*domain.ValidationError {
	return urv.validator.Run(req)
}

func (urv *userRequestValidation) UpdateProfile(req *request.UpdateProfile) []*domain.ValidationError {
	return urv.validator.Run(req)
}
