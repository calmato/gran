package application

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/service"
)

// GroupApplication - GroupApplicationインターフェース
type GroupApplication interface {
	Create(ctx context.Context, req *request.CreateGroup) error
}

type groupApplication struct {
	groupRequestValidation validation.GroupRequestValidation
	groupService           service.GroupService
	userService            service.UserService
}

// NewGroupApplication - GroupApplicationの生成
func NewGroupApplication(
	grv validation.GroupRequestValidation, gs service.GroupService, us service.UserService,
) GroupApplication {
	return &groupApplication{
		groupRequestValidation: grv,
		groupService:           gs,
		userService:            us,
	}
}

func (ga *groupApplication) Create(ctx context.Context, req *request.CreateGroup) error {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if ves := ga.groupRequestValidation.CreateGroup(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	current := time.Now()
	g := &domain.Group{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	if err := ga.groupService.Create(ctx, u, g); err != nil {
		return err
	}

	return nil
}
