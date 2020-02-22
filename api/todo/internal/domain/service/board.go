package service

import (
	"context"
	"time"

	"mime/multipart"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/domain/validation"
)

// BoardService - BoardServiceインターフェース
type BoardService interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
	Create(ctx context.Context, groupID string, b *domain.Board) error
	UploadThumbnail(ctx context.Context, thumbnail multipart.File) (string, error)
}

type boardService struct {
	boardDomainValidation validation.BoardDomainValidation
	boardRepository       repository.BoardRepository
}

// NewBoardService - BoardServiceの生成
func NewBoardService(bdv validation.BoardDomainValidation, br repository.BoardRepository) BoardService {
	return &boardService{
		boardDomainValidation: bdv,
		boardRepository:       br,
	}
}

func (bs *boardService) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	b, err := bs.boardRepository.Index(ctx, groupID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return b, nil
}

func (bs *boardService) Create(ctx context.Context, groupID string, b *domain.Board) error {
	if ves := bs.boardDomainValidation.Board(ctx, b); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()
	b.ID = uuid.New().String()
	b.GroupID = groupID
	b.CreatedAt = current
	b.UpdatedAt = current

	if err := bs.boardRepository.Create(ctx, b); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (bs *boardService) UploadThumbnail(ctx context.Context, thumbnail multipart.File) (string, error) {
	return "", nil
}
