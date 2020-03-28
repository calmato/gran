package validation

import (
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

// BoardRequestValidation - ユーザー関連のバリデーション
type BoardRequestValidation interface {
	CreateBoard(req *request.CreateBoard) []*domain.ValidationError
	CreateBoardList(req *request.CreateBoardList) []*domain.ValidationError
	UpdateBoardList(req *request.UpdateBoardList) []*domain.ValidationError
	UpdateKanban(req *request.UpdateKanban) []*domain.ValidationError
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

func (brv *boardRequestValidation) CreateBoard(req *request.CreateBoard) []*domain.ValidationError {
	return brv.validator.Run(req)
}

func (brv *boardRequestValidation) CreateBoardList(req *request.CreateBoardList) []*domain.ValidationError {
	return brv.validator.Run(req)
}

func (brv *boardRequestValidation) UpdateBoardList(req *request.UpdateBoardList) []*domain.ValidationError {
	return brv.validator.Run(req)
}

func (brv *boardRequestValidation) UpdateKanban(req *request.UpdateKanban) []*domain.ValidationError {
	return brv.validator.Run(req)
}
