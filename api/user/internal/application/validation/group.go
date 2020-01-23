package validation

import "github.com/16francs/gran/api/user/internal/application/request"

// GroupRequestValidation - ユーザー関連のバリデーション
type GroupRequestValidation interface {
	CreateGroup(cg *request.CreateGroup) error
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

func (grv *groupRequestValidation) CreateGroup(cg *request.CreateGroup) error {
	return grv.validator.Run(cg)
}
