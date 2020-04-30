package persistence

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/repository"
	"github.com/calmato/gran/api/todo/lib/firebase/firestore"
)

type groupPersistence struct {
	firestore *firestore.Firestore
}

// NewGroupPersistence - GroupRepositoryの生成
func NewGroupPersistence(fs *firestore.Firestore) repository.GroupRepository {
	return &groupPersistence{
		firestore: fs,
	}
}

func (gp *groupPersistence) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	gs := make([]*domain.Group, len(u.GroupIDs))
	groupCollection := getGroupCollection()

	for i, v := range u.GroupIDs {
		doc, err := gp.firestore.Get(ctx, groupCollection, v)
		if err != nil {
			return nil, err
		}

		g := &domain.Group{}

		err = doc.DataTo(g)
		if err != nil {
			return nil, err
		}

		gs[i] = g
	}

	return gs, nil
}

func (gp *groupPersistence) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	groupCollection := getGroupCollection()

	doc, err := gp.firestore.Get(ctx, groupCollection, groupID)
	if err != nil {
		return nil, err
	}

	g := &domain.Group{}

	err = doc.DataTo(g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (gp *groupPersistence) Create(ctx context.Context, g *domain.Group) error {
	groupCollection := getGroupCollection()

	if err := gp.firestore.Set(ctx, groupCollection, g.ID, g); err != nil {
		return err
	}

	return nil
}

func (gp *groupPersistence) Update(ctx context.Context, g *domain.Group) error {
	groupCollection := getGroupCollection()

	if err := gp.firestore.Set(ctx, groupCollection, g.ID, g); err != nil {
		return err
	}

	return nil
}
