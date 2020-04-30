package validation

import (
	"github.com/calmato/gran/api/todo/internal/application/request"
	"github.com/calmato/gran/api/todo/internal/domain"
)

// GroupRequestValidation - ユーザー関連のバリデーション
type GroupRequestValidation interface {
	CreateGroup(req *request.CreateGroup) []*domain.ValidationError
	UpdateGroup(req *request.UpdateGroup) []*domain.ValidationError
	InviteUsers(req *request.InviteUsers) []*domain.ValidationError
}

type groupRequestValidation struct {
	validator RequestValidator
}

// NewGroupRequestValidation - GroupRequestValidationの生成
func NewGroupRequestValidation() GroupRequestValidation {
	rv := NewRequestValidator()

	return &groupRequestValidation{
		validator: rv,
	}
}

func (grv *groupRequestValidation) CreateGroup(req *request.CreateGroup) []*domain.ValidationError {
	return grv.validator.Run(req)
}

func (grv *groupRequestValidation) UpdateGroup(req *request.UpdateGroup) []*domain.ValidationError {
	return grv.validator.Run(req)
}

func (grv *groupRequestValidation) InviteUsers(req *request.InviteUsers) []*domain.ValidationError {
	return grv.validator.Run(req)
}
