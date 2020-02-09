package validation

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
	dv "github.com/16francs/gran/api/group/internal/domain/validation"
)

type groupDomainValidation struct{}

// NewGroupDomainValidation - GroupDomainValidationの生成
func NewGroupDomainValidation() dv.GroupDomainValidation {
	return &groupDomainValidation{}
}

func (gdv *groupDomainValidation) Group(ctx context.Context, g *domain.Group) []*domain.ValidationError {
	ves := make([]*domain.ValidationError, 0)
	return ves
}
