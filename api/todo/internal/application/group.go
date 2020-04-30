package application

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/calmato/gran/api/todo/internal/application/request"
	"github.com/calmato/gran/api/todo/internal/application/validation"
	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/service"
)

// GroupApplication - GroupApplicationインターフェース
type GroupApplication interface {
	Index(ctx context.Context) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, req *request.CreateGroup) error
	Update(ctx context.Context, groupID string, req *request.UpdateGroup) error
	InviteUsers(ctx context.Context, groupID string, req *request.InviteUsers) error
	Join(ctx context.Context, groupID string) error
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

	if !ga.groupService.IsContainInUserIDs(ctx, u.ID, g) {
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

	g := &domain.Group{
		Name:        req.Name,
		Description: req.Description,
	}

	if _, err := ga.groupService.Create(ctx, u, g); err != nil {
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

	if !ga.groupService.IsContainInUserIDs(ctx, u.ID, g) {
		err = xerrors.New("Failed to Application")
		return domain.Forbidden.New(err)
	}

	if ves := ga.groupRequestValidation.UpdateGroup(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	g.Name = req.Name
	g.Description = req.Description

	if err := ga.groupService.Update(ctx, g); err != nil {
		return err
	}

	return nil
}

func (ga *groupApplication) InviteUsers(ctx context.Context, groupID string, req *request.InviteUsers) error {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	g, err := ga.groupService.Show(ctx, groupID)
	if err != nil {
		return err
	}

	if !ga.groupService.IsContainInUserIDs(ctx, u.ID, g) {
		err = xerrors.New("Failed to Application")
		return domain.Forbidden.New(err)
	}

	if ves := ga.groupRequestValidation.InviteUsers(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	g.InvitedEmails = append(g.InvitedEmails, req.Emails...)

	if err := ga.groupService.InviteUsers(ctx, g); err != nil {
		return err
	}

	return nil
}

func (ga *groupApplication) Join(ctx context.Context, groupID string) error {
	u, err := ga.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	g, err := ga.groupService.Show(ctx, groupID)
	if err != nil {
		return err
	}

	if !ga.groupService.IsContainInInvitedEmails(ctx, u.Email, g) {
		err = xerrors.New("Failed to Application")
		return domain.Forbidden.New(err)
	}

	g.UserIDs = append(g.UserIDs, u.ID)

	removeEmailInInvitedEmails(u.Email, g)

	if err := ga.groupService.Join(ctx, g); err != nil {
		return err
	}

	return nil
}

func removeEmailInInvitedEmails(email string, g *domain.Group) {
	list := make([]string, 0)

	for _, v := range g.InvitedEmails {
		if email == v {
			continue
		}

		list = append(list, v)
	}

	g.InvitedEmails = list
}
