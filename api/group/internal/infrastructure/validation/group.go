package validation

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/repository"
	dv "github.com/16francs/gran/api/group/internal/domain/validation"
)

type groupDomainValidation struct {
	groupRepository repository.GroupRepository
}

// NewGroupDomainValidation - GroupDomainValidationの生成
func NewGroupDomainValidation(gr repository.GroupRepository) dv.GroupDomainValidation {
	return &groupDomainValidation{
		groupRepository: gr,
	}
}

func (gdv *groupDomainValidation) Group(ctx context.Context, g *domain.Group) []*domain.ValidationError {
	return nil
}
