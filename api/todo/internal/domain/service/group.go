package service

import (
	"context"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"golang.org/x/xerrors"
)

// GroupService - GroupServiceインターフェース
type GroupService interface {
	Show(ctx context.Context, groupID string) (*domain.Group, error)
}

type groupService struct {
	groupRepository repository.GroupRepository
}

// NewGroupService - GroupSerivceの生成
func NewGroupService(gr repository.GroupRepository) GroupService {
	return &groupService{
		groupRepository: gr,
	}
}

func (gs *groupService) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g, err := gs.groupRepository.Show(ctx, groupID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		// TODO: エラー処理ちゃんとやる
		return nil, domain.ErrorInDatastore.New(err)
	}
}
