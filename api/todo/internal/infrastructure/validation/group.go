package validation

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	dv "github.com/16francs/gran/api/todo/internal/domain/validation"
)

type groupDomainValidation struct{}

// NewGroupDomainValidation - GroupDomainValidationの生成
func NewGroupDomainValidation() dv.GroupDomainValidation {
	return &groupDomainValidation{}
}

func (gdv *groupDomainValidation) Group(ctx context.Context, g *domain.Group) []*domain.ValidationError {
	ves := make([]*domain.ValidationError, 0)

	if err := uniqueCheckInvitedEmails(g.InvitedEmails); err != nil {
		ve := &domain.ValidationError{
			Field:   "招待メールアドレス一覧",
			Message: dv.UniqueMessage,
		}

		ves = append(ves, ve)
	}

	return ves
}

func uniqueCheckInvitedEmails(invitedEmails []string) error {
	m := make(map[string]struct{})

	for _, v := range invitedEmails {
		if _, ok := m[v]; ok {
			return xerrors.New("There are duplicate values.")
		}

		m[v] = struct{}{}
	}

	return nil
}
