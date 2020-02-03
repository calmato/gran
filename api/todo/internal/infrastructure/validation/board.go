package validation

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
	dv "github.com/16francs/gran/api/todo/internal/domain/validation"
)

type boardDomainValidation struct {
	validator DomainValidator
}

// NewGroupDomainValidation - GroupDomainValidationの生成
func NewBoardDomainValidation() dv.BoardDomainValidation {
	v := NewDomainValidator()

	return &boardDomainValidation{
		validator: v,
	}
}

func (bdv *boardDomainValidation) Board(ctx context.Context, b *domain.Board) []*domain.ValidationError {
	return bdv.validator.Run(b)
}
