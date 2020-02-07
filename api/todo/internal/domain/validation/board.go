package validation

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
)

// BoardDomainValidation - BoardDomainValidationインターフェース
type BoardDomainValidation interface {
	Board(ctx context.Context, b *domain.Board) []*domain.ValidationError
}
