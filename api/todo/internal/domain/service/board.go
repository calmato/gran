package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/domain/validation"
)

// BoardService - BoardServiceインターフェース
type BoardService interface {
	Create(ctx context.Context, groupID string, b *domain.Board) error
}

type boardService struct {
	BoardDomainValidation validation.BoardDomainValidation
	BoardRepository       repository.BoardRepository
}

// NewBoardService - BoardServiceの生成
func NewBoardService(bdv validation.BoardDomainValidation, br repository.BoardRepository) BoardService {
	return &boardService{
		BoardDomainValidation: bdv,
		BoardRepository:       br,
	}
}

func (bs *boardService) Create(ctx context.Context, groupID string, b *domain.Board) error {
	if ves := bs.BoardDomainValidation.Board(ctx, b); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := bs.BoardRepository.Create(ctx, groupID, b); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}
