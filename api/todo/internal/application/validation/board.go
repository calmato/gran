package validation

import (
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

// BoardRequestValidation - ユーザー関連のバリデーション
type BoardRequestValidation interface {
	CreateBoard(cb *request.CreateBoard) []*domain.ValidationError
}

type boardRequestValidation struct {
	validator RequestValidator
}

// NewBoardRequestValidation - BoardRequestValidationの生成
func NewBoardRequestValidation() BoardRequestValidation {
	rv := NewRequestValidator()

	return &boardRequestValidation{
		validator: rv,
	}
}

func (brv *boardRequestValidation) CreateBoard(cb *request.CreateBoard) []*domain.ValidationError {
	return brv.validator.Run(cb)
}
