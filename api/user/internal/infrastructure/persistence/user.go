package persistence

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type userPersistence struct {
	firestore firestore.Firestore
}

// UserCollection - UserCollection名
const UserCollection = "users"

// NewUserPersistence - UserRepositoryの生成
func NewUserPersistence(fs firestore.Firestore) repository.UserRepository {
	return &userPersistence{
		firestore: fs,
	}
}

func (r *userPersistence) Create(ctx context.Context, u *domain.User) error {
	// TODO: addInFirebaseの実装

	if err := addInFirestore(ctx, r.firestore, u); err != nil {
		return xerrors.Errorf("Failed to UserPersistence/Create: %w", err)
	}

	return nil
}

func addInFirestore(ctx context.Context, fs firestore.Firestore, u *domain.User) error {
	return fs.Add(ctx, UserCollection, u)
}
