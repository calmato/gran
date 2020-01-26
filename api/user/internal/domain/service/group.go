package service

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain/repository"
	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/validation"
)

// GroupService - GroupServiceインターフェース
type GroupService interface {
	Create(ctx context.Context, u *domain.User, g *domain.Group) error
}

type groupService struct {
	GroupDomainValidation validation.GroupDomainValidation
	GroupRepository       repository.GroupRepository
}

// NewGroupService - GroupServiceの生成
func NewGroupService(gdv validation.GroupDomainValidation, gr repository.GroupRepository) GroupService {
	return &groupService{
		GroupDomainValidation: gdv,
		GroupRepository:       gr,
	}
}

func (gs *groupService) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	if err := gs.GroupDomainValidation.Group(ctx, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/DomainValidation: %w", err)
		return domain.InvalidDomainValidation.New(err)
	}

	if err := gs.GroupRepository.Create(ctx, u, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.Unknown.New(err)
	}

	return nil
}
