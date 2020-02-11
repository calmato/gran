package application

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/service"
)

// BoardApplication - BoardApplicationインターフェース
type BoardApplication interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
	Create(ctx context.Context, req *request.CreateBoard) error
}

type boardApplication struct {
	boardRequestValidation validation.BoardRequestValidation
	boardService           service.BoardService
	userService            service.UserService
}

// NewBoardApplication - BoardApplicationの生成
func NewBoardApplication(
	brv validation.BoardRequestValidation, bs service.BoardService, us service.UserService,
) BoardApplication {
	return &boardApplication{
		boardRequestValidation: brv,
		boardService:           bs,
		userService:            us,
	}
}

func (ba *boardApplication) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to create Board in the Group")
		return nil, domain.Forbidden.New(err)
	}

	bs, err := ba.boardService.Index(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (ba *boardApplication) Create(ctx context.Context, req *request.CreateBoard) error {
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if ves := ba.boardRequestValidation.CreateBoard(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	b := &domain.Board{
		Name:            req.Name,
		Closed:          req.Closed,
		ThumbnailURL:    req.ThumbnailURL,
		BackgroundColor: req.BackgroundColor,
		Labels:          req.Labels,
	}

	groupID := req.GroupID

	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to create Board in the Group")
		return domain.Forbidden.New(err)
	}

	if err := ba.boardService.Create(ctx, groupID, b); err != nil {
		return err
	}

	return nil
}
