package validation

import (
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

// TaskRequestValidation - ユーザー関連のバリデーション
type TaskRequestValidation interface {
	CreateTask(req *request.CreateTask) []*domain.ValidationError
}

type taskRequestValidation struct {
	validator RequestValidator
}

// NewTaskRequestValidation - TaskRequestValidationの生成
func NewTaskRequestValidation() TaskRequestValidation {
	rv := NewRequestValidator()

	return &taskRequestValidation{
		validator: rv,
	}
}

func (brv *taskRequestValidation) CreateTask(req *request.CreateTask) []*domain.ValidationError {
	return brv.validator.Run(req)
}
