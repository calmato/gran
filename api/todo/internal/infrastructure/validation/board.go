package validation

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
	dv "github.com/calmato/gran/api/todo/internal/domain/validation"
)

type boardDomainValidation struct{}

// NewBoardDomainValidation - BoardDomainValidationの生成
func NewBoardDomainValidation() dv.BoardDomainValidation {
	return &boardDomainValidation{}
}

func (bdv *boardDomainValidation) Board(ctx context.Context, b *domain.Board) []*domain.ValidationError {
	ves := make([]*domain.ValidationError, 0)

	return ves
}

func (bdv *boardDomainValidation) BoardList(ctx context.Context, b *domain.BoardList) []*domain.ValidationError {
	ves := make([]*domain.ValidationError, 0)

	return ves
}
