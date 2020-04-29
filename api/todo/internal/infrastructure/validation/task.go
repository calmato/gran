package validation

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
	dv "github.com/calmato/gran/api/todo/internal/domain/validation"
)

type taskDomainValidation struct{}

// NewTaskDomainValidation - TaskDomainValidationの生成
func NewTaskDomainValidation() dv.TaskDomainValidation {
	return &taskDomainValidation{}
}

func (tdv *taskDomainValidation) Task(ctx context.Context, b *domain.Task) []*domain.ValidationError {
	ves := make([]*domain.ValidationError, 0)

	return ves
}
