package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/domain/uploader"
	"github.com/16francs/gran/api/todo/internal/domain/validation"
)

// BoardService - BoardServiceインターフェース
type BoardService interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
	Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error)
	Create(ctx context.Context, b *domain.Board) (*domain.Board, error)
	UploadThumbnail(ctx context.Context, data []byte) (string, error)
	ShowBoardList(ctx context.Context, groupID string, boardID string, boardListID string) (*domain.BoardList, error)
	CreateBoardList(ctx context.Context, groupID string, boardID string, bl *domain.BoardList) (*domain.BoardList, error)
	UpdateBoardList(ctx context.Context, groupID string, boardID string, bl *domain.BoardList) (*domain.BoardList, error)
	UpdateKanban(ctx context.Context, groupID string, boardID string, b *domain.Board) error
}

type boardService struct {
	boardDomainValidation validation.BoardDomainValidation
	boardRepository       repository.BoardRepository
	taskRepository        repository.TaskRepository
	fileUploader          uploader.FileUploader
}

// NewBoardService - BoardServiceの生成
func NewBoardService(
	bdv validation.BoardDomainValidation, br repository.BoardRepository,
	tr repository.TaskRepository, fu uploader.FileUploader,
) BoardService {
	return &boardService{
		boardDomainValidation: bdv,
		boardRepository:       br,
		taskRepository:        tr,
		fileUploader:          fu,
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

func (bs *boardService) Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error) {
	b, err := bs.boardRepository.Show(ctx, groupID, boardID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	bls, err := bs.boardRepository.IndexBoardList(ctx, groupID, boardID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	b.Lists = make(map[string]*domain.BoardList)
	for _, bl := range bls {
		b.Lists[bl.ID] = bl

		ts, err := bs.taskRepository.IndexByBoardListID(ctx, bl.ID)
		if err != nil {
			err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
			return nil, domain.ErrorInDatastore.New(err)
		}

		bl.Tasks = make(map[string]*domain.Task)
		for _, t := range ts {
			bl.Tasks[t.ID] = t
		}
	}

	return b, nil
}

func (bs *boardService) Create(ctx context.Context, b *domain.Board) (*domain.Board, error) {
	if ves := bs.boardDomainValidation.Board(ctx, b); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()

	b.ID = uuid.New().String()
	b.CreatedAt = current
	b.UpdatedAt = current

	if err := bs.boardRepository.Create(ctx, b); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return b, nil
}

func (bs *boardService) UploadThumbnail(ctx context.Context, data []byte) (string, error) {
	thumbnailURL, err := bs.fileUploader.UploadBoardThumbnail(ctx, data)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Uploader: %w", err)
		return "", domain.ErrorInStorage.New(err)
	}

	return thumbnailURL, nil
}

func (bs *boardService) ShowBoardList(
	ctx context.Context, groupID string, boardID string, boardListID string,
) (*domain.BoardList, error) {
	bl, err := bs.boardRepository.ShowBoardList(ctx, groupID, boardID, boardListID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return bl, nil
}

func (bs *boardService) CreateBoardList(
	ctx context.Context, groupID string, boardID string, bl *domain.BoardList,
) (*domain.BoardList, error) {
	if ves := bs.boardDomainValidation.BoardList(ctx, bl); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()

	bl.ID = uuid.New().String()
	bl.CreatedAt = current
	bl.UpdatedAt = current

	if err := bs.boardRepository.CreateBoardList(ctx, groupID, boardID, bl); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	b, err := bs.boardRepository.Show(ctx, groupID, boardID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	b.ListIDs = append(b.ListIDs, bl.ID)

	if err := bs.boardRepository.Update(ctx, b); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return bl, nil
}

func (bs *boardService) UpdateBoardList(
	ctx context.Context, groupID string, boardID string, bl *domain.BoardList,
) (*domain.BoardList, error) {
	if ves := bs.boardDomainValidation.BoardList(ctx, bl); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()

	bl.UpdatedAt = current

	if err := bs.boardRepository.UpdateBoardList(ctx, groupID, boardID, bl); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return bl, nil
}

// UpdateKanban - ボードリスト, タスク順序の編集
func (bs *boardService) UpdateKanban(ctx context.Context, groupID string, boardID string, b *domain.Board) error {
	if ves := bs.boardDomainValidation.Board(ctx, b); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := bs.boardRepository.Update(ctx, b); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	for _, bl := range b.Lists {
		if ves := bs.boardDomainValidation.BoardList(ctx, bl); len(ves) > 0 {
			err := xerrors.New("Failed to Domain/DomainValidation")
			return domain.InvalidDomainValidation.New(err, ves...)
		}

		if err := bs.boardRepository.UpdateBoardList(ctx, groupID, boardID, bl); err != nil {
			err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
			return domain.ErrorInDatastore.New(err)
		}
	}

	return nil
}
