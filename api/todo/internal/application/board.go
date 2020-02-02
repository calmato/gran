package application

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/service"
)

// BoardApplication - BoardApplicationインターフェース
type BoardApplication interface {
	Create(ctx context.Context, req *request.CreateBoard) error
}

type boardApplication struct {
	BoardRequestValidation validation.BoardRequestValidation
	BoardService           service.BoardService
	UserService            service.UserService
}

// NewBoardApplication - BoardApplicationの生成
func NewBoardApplication(
	brv validation.BoardRequestValidation, bs service.BoardService, us service.UserService,
) BoardApplication {
	return &boardApplication{
		BoardRequestValidation: brv,
		BoardService:           bs,
		UserService:            us,
	}
}

func (ba *boardApplication) Create(ctx context.Context, req *request.CreateBoard) error {
	_, err := ba.UserService.Authentication(ctx)
	if err != nil {
		return err
	}

	if ves := ba.BoardRequestValidation.CreateBoard(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	current := time.Now()
	b := &domain.Board{
		Name:            req.Name,
		Closed:          req.Closed,
		ThumbnailURL:    req.ThumbnailURL,
		BackgroundColor: req.BackgroundColor,
		Labels:          req.Labels,
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	groupID := req.GroupID

	if err := ba.BoardService.Create(ctx, groupID, b); err != nil {
		return err
	}

	return nil
}
