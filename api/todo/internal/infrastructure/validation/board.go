package validation

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
	dv "github.com/16francs/gran/api/todo/internal/domain/validation"
)

type boardDomainValidation struct{}

// NewBoardDomainValidation - GroupDomainValidationの生成
func NewBoardDomainValidation() dv.BoardDomainValidation {
	return &boardDomainValidation{}
}

func (bdv *boardDomainValidation) Board(ctx context.Context, b *domain.Board) []*domain.ValidationError {
	return nil
}
