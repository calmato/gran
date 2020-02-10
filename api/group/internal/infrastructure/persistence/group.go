package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/repository"
	"github.com/16francs/gran/api/group/lib/firebase/firestore"
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
	gs := make([]*domain.Group, len(u.GroupRefs))

	for i, v := range u.GroupRefs {
		doc, err := gp.firestore.Get(ctx, GroupCollection, GetGroupID(v))
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
	doc, err := gp.firestore.Get(ctx, GroupCollection, groupID)
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

func (gp *groupPersistence) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	g.ID = uuid.New().String()
	g.UserRefs = append(g.UserRefs, GetUserReference(u.ID))

	if err := gp.firestore.Set(ctx, GroupCollection, g.ID, g); err != nil {
		return err
	}

	current := time.Now()

	u.GroupRefs = append(u.GroupRefs, GetGroupReference(g.ID))
	u.UpdatedAt = current

	if err := gp.firestore.Set(ctx, UserCollection, u.ID, u); err != nil {
		return err
	}

	return nil
}

func (gp *groupPersistence) Update(ctx context.Context, g *domain.Group) error {
	if err := gp.firestore.Set(ctx, GroupCollection, g.ID, g); err != nil {
		return err
	}

	return nil
}

func (gp *groupPersistence) UserIDExistsInUserRefs(ctx context.Context, userID string, g *domain.Group) bool {
	userRef := GetUserReference(userID)

	for _, v := range g.UserRefs {
		if userRef == v {
			return true
		}
	}

	return false
}

func (gp *groupPersistence) EmailExistsInInvitedEmails(ctx context.Context, email string, g *domain.Group) bool {
	for _, v := range g.InvitedEmails {
		if email == v {
			return true
		}
	}

	return false
}
