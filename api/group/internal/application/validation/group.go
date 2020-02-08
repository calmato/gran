package validation

import (
	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
)

// GroupRequestValidation - ユーザー関連のバリデーション
type GroupRequestValidation interface {
	CreateGroup(cg *request.CreateGroup) []*domain.ValidationError
	UpdateGroup(ug *request.UpdateGroup) []*domain.ValidationError
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

func (grv *groupRequestValidation) CreateGroup(cg *request.CreateGroup) []*domain.ValidationError {
	return grv.validator.Run(cg)
}

func (grv *groupRequestValidation) UpdateGroup(ug *request.UpdateGroup) []*domain.ValidationError {
	return grv.validator.Run(ug)
}
