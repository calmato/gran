package validation

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	dv "github.com/16francs/gran/api/user/internal/domain/validation"
)

type groupDomainValidation struct {
	validator       DomainValidator
	groupRepository repository.GroupRepository
}

// NewGroupDomainValidation - GroupDomainValidationの生成
func NewGroupDomainValidation(gr repository.GroupRepository) dv.GroupDomainValidation {
	v := NewDomainValidator()

	return &groupDomainValidation{
		validator:       v,
		groupRepository: gr,
	}
}

func (gdv *groupDomainValidation) Group(ctx context.Context, g *domain.Group) error {
	return gdv.validator.Run(g)
}
