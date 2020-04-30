package validation

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
)

// TaskDomainValidation - TaskDomainValidationインターフェース
type TaskDomainValidation interface {
	Task(ctx context.Context, t *domain.Task) []*domain.ValidationError
}
