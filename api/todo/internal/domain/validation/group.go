package validation

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
)

// GroupDomainValidation - GroupDomainRepositoryインターフェース
type GroupDomainValidation interface {
	Group(ctx context.Context, g *domain.Group) []*domain.ValidationError
}
