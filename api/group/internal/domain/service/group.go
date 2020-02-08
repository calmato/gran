package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/repository"
	"github.com/16francs/gran/api/group/internal/domain/validation"
)

// GroupService - GroupServiceインターフェース
type GroupService interface {
	Index(ctx context.Context, u *domain.User) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, u *domain.User, g *domain.Group) error
	Update(ctx context.Context, g *domain.Group) error
	InviteUser(ctx context.Context, userID string, g *domain.Group) error
	ExistUserIDInUserRefs(ctx context.Context, userID string, g *domain.Group) bool
	ExistUserIDInInvitedUserRefs(ctx context.Context, userID string, g *domain.Group) bool
}

type groupService struct {
	groupDomainValidation validation.GroupDomainValidation
	groupRepository       repository.GroupRepository
}

// NewGroupService - GroupServiceの生成
func NewGroupService(gdv validation.GroupDomainValidation, gr repository.GroupRepository) GroupService {
	return &groupService{
		groupDomainValidation: gdv,
		groupRepository:       gr,
	}
}

func (gs *groupService) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	g, err := gs.groupRepository.Index(ctx, u)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return g, nil
}

func (gs *groupService) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g, err := gs.groupRepository.Show(ctx, groupID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return g, nil
}

func (gs *groupService) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	if ves := gs.groupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := gs.groupRepository.Create(ctx, u, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (gs *groupService) Update(ctx context.Context, g *domain.Group) error {
	if ves := gs.groupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := gs.groupRepository.Update(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (gs *groupService) InviteUser(ctx context.Context, userID string, g *domain.Group) error {
	if err := gs.groupRepository.InviteUser(ctx, userID, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (gs *groupService) ExistUserIDInUserRefs(ctx context.Context, userID string, g *domain.Group) bool {
	return gs.groupRepository.ExistUserIDInUserRefs(ctx, userID, g)
}

func (gs *groupService) ExistUserIDInInvitedUserRefs(ctx context.Context, userID string, g *domain.Group) bool {
	return gs.groupRepository.ExistUserIDInInvitedUserRefs(ctx, userID, g)
}
