package persistence

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type userPersistence struct {
	auth      authentication.Auth
	firestore firestore.Firestore
}

// UserCollection - UserCollection名
const UserCollection = "users"

// NewUserPersistence - UserRepositoryの生成
func NewUserPersistence(fa authentication.Auth, fs firestore.Firestore) repository.UserRepository {
	return &userPersistence{
		auth:      fa,
		firestore: fs,
	}
}

func (r *userPersistence) Create(ctx context.Context, u *domain.User) error {
	uid, err := createUserInAuth(ctx, r.auth, u)
	if err != nil {
		return xerrors.Errorf("Failed to UserPersistence/Create: %w", err)
	}

	if err = addInFirestore(ctx, r.firestore, u); err != nil {
		return xerrors.Errorf("Failed to UserPersistence/Create: %w", err)
	}

	return nil
}

func createUserInAuth(ctx context.Context, fa authentication.Auth, u *domain.User) (string, error) {
	return fa.CreateUser(ctx, u.Email, u.Password)
}

func addInFirestore(ctx context.Context, fs firestore.Firestore, u *domain.User) error {
	return fs.Add(ctx, UserCollection, u)
}
