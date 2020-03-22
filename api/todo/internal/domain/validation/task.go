package validation

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// TaskDomainValidation - TaskDomainValidationインターフェース
type TaskDomainValidation interface {
	Task(ctx context.Context, t *domain.Task) []*domain.ValidationError
}
