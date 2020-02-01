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
	Create(ctx context.Context, req *request.CreateGroup) error
}

type groupApplication struct {
	GroupRequestValidation validation.GroupRequestValidation
	GroupService           service.GroupService
	UserService            service.UserService
}

// NewGroupApplication - GroupApplicationの生成
func NewGroupApplication(
	grv validation.GroupRequestValidation, gs service.GroupService, us service.UserService,
) GroupApplication {
	return &groupApplication{
		GroupRequestValidation: grv,
		GroupService:           gs,
		UserService:            us,
	}
}

func (ga *groupApplication) Index(ctx context.Context) ([]*domain.Group, error) {
	u, err := ga.UserService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	gs, err := ga.GroupService.Index(ctx, u)
	if err != nil {
		return nil, err
	}

	return gs, nil
}

func (ga *groupApplication) Create(ctx context.Context, req *request.CreateGroup) error {
	u, err := ga.UserService.Authentication(ctx)
	if err != nil {
		return err
	}

	if ves := ga.GroupRequestValidation.CreateGroup(req); len(ves) > 0 {
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

	if err := ga.GroupService.Create(ctx, u, g); err != nil {
		return err
	}

	return nil
}
