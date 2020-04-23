package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/domain/validation"
)

// GroupService - GroupServiceインターフェース
type GroupService interface {
	Index(ctx context.Context, u *domain.User) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, u *domain.User, g *domain.Group) (*domain.Group, error)
	Update(ctx context.Context, g *domain.Group) error
	InviteUsers(ctx context.Context, g *domain.Group) error
	Join(ctx context.Context, g *domain.Group) error
	IsContainInUserIDs(ctx context.Context, userID string, g *domain.Group) bool
	IsContainInInvitedEmails(ctx context.Context, email string, g *domain.Group) bool
}

type groupService struct {
	groupDomainValidation validation.GroupDomainValidation
	groupRepository       repository.GroupRepository
	userRepository        repository.UserRepository
	boardRepository       repository.BoardRepository
}

// NewGroupService - GroupServiceの生成
func NewGroupService(
	gdv validation.GroupDomainValidation, gr repository.GroupRepository,
	ur repository.UserRepository, br repository.BoardRepository,
) GroupService {
	return &groupService{
		groupDomainValidation: gdv,
		groupRepository:       gr,
		userRepository:        ur,
		boardRepository:       br,
	}
}

func (gs *groupService) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	groups, err := gs.groupRepository.Index(ctx, u)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	for _, g := range groups {
		bs, err := gs.boardRepository.Index(ctx, g.ID)
		if err != nil {
			err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
			return nil, domain.ErrorInDatastore.New(err)
		}

		boardIDs := make([]string, len(bs))
		for i, b := range bs {
			boardIDs[i] = b.ID
		}

		g.BoardIDs = boardIDs
	}

	return groups, nil
}

func (gs *groupService) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g, err := gs.groupRepository.Show(ctx, groupID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return g, nil
}

func (gs *groupService) Create(ctx context.Context, u *domain.User, g *domain.Group) (*domain.Group, error) {
	if ves := gs.groupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return nil, domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()
	g.ID = uuid.New().String()
	g.UserIDs = append(g.UserIDs, u.ID)
	g.CreatedAt = current
	g.UpdatedAt = current

	if err := gs.groupRepository.Create(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	u.GroupIDs = append(u.GroupIDs, g.ID)
	u.UpdatedAt = current

	if err := gs.userRepository.Update(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return g, nil
}

func (gs *groupService) Update(ctx context.Context, g *domain.Group) error {
	if ves := gs.groupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	current := time.Now()
	g.UpdatedAt = current

	if err := gs.groupRepository.Update(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (gs *groupService) InviteUsers(ctx context.Context, g *domain.Group) error {
	if ves := gs.groupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")

		if isContainUniqueError(ves) {
			return domain.InvalidDomainValidation.New(err, ves...)
		}

		return domain.Unknown.New(err, ves...)
	}

	if err := gs.groupRepository.Update(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (gs *groupService) Join(ctx context.Context, g *domain.Group) error {
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

func (gs *groupService) IsContainInUserIDs(ctx context.Context, userID string, g *domain.Group) bool {
	for _, v := range g.UserIDs {
		if userID == v {
			return true
		}
	}

	return false
}

func (gs *groupService) IsContainInInvitedEmails(ctx context.Context, email string, g *domain.Group) bool {
	for _, v := range g.InvitedEmails {
		if email == v {
			return true
		}
	}

	return false
}

func isContainUniqueError(ves []*domain.ValidationError) bool {
	for _, v := range ves {
		if v.Message == validation.UniqueMessage {
			return true
		}
	}

	return false
}
