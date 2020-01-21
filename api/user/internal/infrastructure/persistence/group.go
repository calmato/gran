package persistence

import (
	"context"
	"strings"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type groupPersistence struct {
	auth      *authentication.Auth
	firestore *firestore.Firestore
}

func (gp *groupPersistence) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	g.UserRefs = append(g.UserRefs, getUserCollection(u.ID))

	return gp.firestore.Add(ctx, GroupCollection, g)
}

func getUserCollection(userID string) string {
	return strings.Join([]string{UserCollection, userID}, "/")
}
