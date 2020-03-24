package application

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/service"
	"golang.org/x/xerrors"
)

// TaskApplication - TaskApplicationインターフェース
type TaskApplication interface {
	Create(ctx context.Context, groupID string, boardID string, req *request.CreateTask) error
}

type taskApplication struct {
	taskRequestValidation validation.TaskRequestValidation
	taskService           service.TaskService
	userService           service.UserService
}

// NewTaskApplication - TaskApplicationの生成
func NewTaskApplication(
	trv validation.TaskRequestValidation, ts service.TaskService, us service.UserService,
) TaskApplication {
	return &taskApplication{
		taskRequestValidation: trv,
		taskService:           ts,
		userService:           us,
	}
}

func (ta *taskApplication) Create(ctx context.Context, groupID string, boardID string, req *request.CreateTask) error {
	u, err := ta.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if !ta.userService.IsContainInGroupIDs(ctx, groupID, u) {
		err := xerrors.New("Unable to create Board in the Group")
		return domain.Forbidden.New(err)
	}

	if ves := ta.taskRequestValidation.CreateTask(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	// TODO: ボード、ボードリストが存在するかの検証

	attachmentURLs := make([]string, len(req.Attachments))

	for i, v := range req.Attachments {
		// data:image/png;base64,iVBORw0KGgoAAAA... みたいなのうちの
		// `data:image/png;base64,` の部分を無くした []byte を取得
		b64data := v[strings.IndexByte(v, ',')+1:]

		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			err = xerrors.Errorf("Failed to Application: %w", err)
			return domain.Unknown.New(err)
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
		BoardID:         boardID,
		BoardListID:     req.BoardListID,
	}

	if _, err := ta.taskService.Create(ctx, groupID, boardID, t); err != nil {
		return err
	}

	return nil
}
