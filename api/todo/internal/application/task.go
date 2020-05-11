package application

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/calmato/gran/api/todo/internal/application/request"
	"github.com/calmato/gran/api/todo/internal/application/validation"
	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/service"
	"golang.org/x/xerrors"
)

// TaskApplication - TaskApplicationインターフェース
type TaskApplication interface {
	Show(ctx context.Context, taskID string) (*domain.Task, error)
	Create(ctx context.Context, groupID string, boardID string, req *request.CreateTask) (*domain.Task, error)
}

type taskApplication struct {
	taskRequestValidation validation.TaskRequestValidation
	taskService           service.TaskService
	boardService          service.BoardService
	userService           service.UserService
}

// NewTaskApplication - TaskApplicationの生成
func NewTaskApplication(
	trv validation.TaskRequestValidation, ts service.TaskService, bs service.BoardService, us service.UserService,
) TaskApplication {
	return &taskApplication{
		taskRequestValidation: trv,
		taskService:           ts,
		boardService:          bs,
		userService:           us,
	}
}

func (ta *taskApplication) Show(ctx context.Context, taskID string) (*domain.Task, error) {
	u, err := ta.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	t, err := ta.taskService.Show(ctx, taskID)
	if err != nil {
		return nil, domain.NotFound.New(err)
	}

	if !ta.userService.IsContainInGroupIDs(ctx, t.GroupID, u) {
		err := xerrors.New("Unable to create Board in the Group")
		return nil, domain.Forbidden.New(err)
	}

	return t, nil
}

func (ta *taskApplication) Create(
	ctx context.Context, groupID string, boardID string, req *request.CreateTask,
) (*domain.Task, error) {
	u, err := ta.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	if !ta.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to create Board in the Group")
		return nil, domain.Forbidden.New(err)
	}

	if ves := ta.taskRequestValidation.CreateTask(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return nil, domain.InvalidRequestValidation.New(err, ves...)
	}

	if !ta.boardService.ExistsBoardList(ctx, groupID, boardID, req.BoardListID) {
		err := xerrors.New("Unable to create Board in the Group")
		return nil, domain.Forbidden.New(err)
	}

	attachmentURLs := make([]string, len(req.Attachments))

	for i, v := range req.Attachments {
		// data:image/png;base64,iVBORw0KGgoAAAA... みたいなのうちの
		// `data:image/png;base64,` の部分を無くした []byte を取得
		b64data := v[strings.IndexByte(v, ',')+1:]

		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			err = xerrors.Errorf("Failed to Application: %w", err)
			return nil, domain.Unknown.New(err)
		}

		// TODO: ファイルの保存先URL取得
		fmt.Println(i, data)
	}

	t := &domain.Task{
		Name:            req.Name,
		Description:     req.Description,
		Labels:          req.Labels,
		AssignedUserIDs: req.AssignedUserIDs,
		DeadlinedAt:     req.DeadlinedAt,
		AttachmentURLs:  attachmentURLs,
		GroupID:         groupID,
		BoardID:         boardID,
	}

	task, err := ta.taskService.Create(ctx, groupID, boardID, req.BoardListID, t)
	if err != nil {
		return nil, err
	}

	return task, nil
}
