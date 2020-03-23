package application

import (
	"context"
	"encoding/base64"
	"strings"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/service"
)

// BoardApplication - BoardApplicationインターフェース
type BoardApplication interface {
	Index(ctx context.Context, groupID string) ([]*domain.Board, error)
	Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error)
	Create(ctx context.Context, groupID string, req *request.CreateBoard) error
	CreateBoardList(ctx context.Context, groupID string, boardID string, req *request.CreateBoardList) error
	UpdateKanban(ctx context.Context, groupID string, boardID string, req *request.UpdateKanban) error
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
		err := xerrors.New("Unable to IndexBoard function")
		return nil, domain.Forbidden.New(err)
	}

	bs, err := ba.boardService.Index(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (ba *boardApplication) Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error) {
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to ShowBoard function")
		return nil, domain.Forbidden.New(err)
	}

	b, err := ba.boardService.Show(ctx, groupID, boardID)
	if err != nil {
		return nil, err
	}

	return b, err
}

func (ba *boardApplication) Create(ctx context.Context, groupID string, req *request.CreateBoard) error {
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to CreateBoard function")
		return domain.Forbidden.New(err)
	}

	if ves := ba.boardRequestValidation.CreateBoard(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	thumbnailURL := ""

	if req.Thumbnail != "" {
		// data:image/png;base64,iVBORw0KGgoAAAA... みたいなのうちの
		// `data:image/png;base64,` の部分を無くした []byte を取得
		b64data := req.Thumbnail[strings.IndexByte(req.Thumbnail, ',')+1:]

		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			err = xerrors.Errorf("Failed to Application: %w", err)
			return domain.Unknown.New(err)
		}

		thumbnailURL, err = ba.boardService.UploadThumbnail(ctx, data)
		if err != nil {
			return err
		}
	}

	b := &domain.Board{
		Name:            req.Name,
		IsClosed:        req.IsClosed,
		ThumbnailURL:    thumbnailURL,
		BackgroundColor: req.BackgroundColor,
		Labels:          req.Labels,
		GroupID:         groupID,
	}

	if _, err := ba.boardService.Create(ctx, b); err != nil {
		return err
	}

	return nil
}

func (ba *boardApplication) CreateBoardList(
	ctx context.Context, groupID string, boardID string, req *request.CreateBoardList,
) error {
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to CreateBoardList function")
		return domain.Forbidden.New(err)
	}

	// TODO: ボードが存在するかの検証

	if ves := ba.boardRequestValidation.CreateBoardList(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	bl := &domain.BoardList{
		Name:    req.Name,
		Color:   req.Color,
		BoardID: boardID,
		TaskIDs: []string{},
	}

	if _, err := ba.boardService.CreateBoardList(ctx, groupID, boardID, bl); err != nil {
		return err
	}

	return nil
}

// UpdateKanban - ボードリスト, タスク順序の編集
func (ba *boardApplication) UpdateKanban(
	ctx context.Context, groupID string, boardID string, req *request.UpdateKanban,
) error {
	// 認証
	u, err := ba.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	// 認可
	if !ba.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to UpdateKanban function")
		return domain.Forbidden.New(err)
	}

	// 存在性検証
	b, err := ba.boardService.Show(ctx, groupID, boardID)
	if err != nil {
		return err
	}

	// バリデーション
	if ves := ba.boardRequestValidation.UpdateKanban(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	// ListとTaskの数が一致すれば、Datastoreの値と順序以外は一致すると考える
	// List数の検証 -> 一致しなければ、400
	if len(req.Lists) != len(b.ListIDs) {
		err := xerrors.New("Not equal BoardLists length")
		return domain.NotEqualRequestWithDatastore.New(err)
	}

	// Task数の検証 -> 一致しなければ、400
	lenReqTasks := 0
	for _, v := range req.Lists {
		lenReqTasks += len(v.Tasks)
	}

	lenDBTasks := 0
	for _, v := range b.Lists {
		lenDBTasks += len(v.Tasks)
	}

	if lenReqTasks != lenDBTasks {
		err := xerrors.New("Not equal Tasks length")
		return domain.NotEqualRequestWithDatastore.New(err)
	}

	// Board.ListIDs, BoardList.TasksIDの更新
	listIDs := make([]string, len(req.Lists))
	for i, bl := range req.Lists {
		listIDs[i] = bl.ID

		taskIDs := make([]string, len(bl.Tasks))
		for j, t := range bl.Tasks {
			taskIDs[j] = t.ID
		}

		b.Lists[bl.ID].TaskIDs = taskIDs
	}

	b.ListIDs = listIDs

	if err := ba.boardService.UpdateKanban(ctx, groupID, boardID, b); err != nil {
		return err
	}

	return nil
}
