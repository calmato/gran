package application

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/application/validation"
	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/service"
)

// GroupApplication - GroupApplicationインターフェース
type GroupApplication interface {
	Index(ctx context.Context) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, req *request.CreateGroup) error
	Update(ctx context.Context, groupID string, req *request.UpdateGroup) error
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

func (ga *groupApplication) Index(ctx context.Context) ([]*domain.Group, error) {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	gs, err := ga.groupService.Index(ctx, u)
	if err != nil {
		return nil, err
	}

	return gs, nil
}

func (ga *groupApplication) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	g, err := ga.groupService.Show(ctx, groupID)
	if err != nil {
		return nil, err
	}

	if !ga.groupService.UserIDExistsInUserRefs(ctx, u.ID, g) {
		err = xerrors.New("Failed to Application")
		return nil, domain.Forbidden.New(err)
	}

	return g, nil
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

func (ga *groupApplication) Update(ctx context.Context, groupID string, req *request.UpdateGroup) error {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	g, err := ga.groupService.Show(ctx, groupID)
	if err != nil {
		return err
	}

	if !ga.groupService.UserIDExistsInUserRefs(ctx, u.ID, g) {
		err = xerrors.New("Failed to Application")
		return domain.Forbidden.New(err)
	}

	if ves := ga.groupRequestValidation.UpdateGroup(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	current := time.Now()
	g.Name = req.Name
	g.Description = req.Description
	g.UpdatedAt = current

	if err := ga.groupService.Update(ctx, g); err != nil {
		return err
	}

	return nil
}
