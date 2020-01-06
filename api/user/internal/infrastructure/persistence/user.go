package persistence

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type userPersistence struct {
	auth      *authentication.Auth
	firestore *firestore.Firestore
}

// UserCollection - UserCollection名
const UserCollection = "users"

// NewUserPersistence - UserRepositoryの生成
func NewUserPersistence(fa *authentication.Auth, fs *firestore.Firestore) repository.UserRepository {
	return &userPersistence{
		auth:      fa,
		firestore: fs,
	}
}

func (r *userPersistence) Create(ctx context.Context, u *domain.User) error {
	uid, err := createUserInAuth(ctx, r.auth, u)
	if err != nil {
		return err
	}

	u.ID = uid

	if err = setInFirestore(ctx, r.firestore, u); err != nil {
		return err
	}

	return nil
}

func (r *userPersistence) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := getUIDByEmailInAuth(ctx, r.auth, email)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func getUIDByEmailInAuth(ctx context.Context, fa *authentication.Auth, email string) (string, error) {
	return fa.GetUIDByEmail(ctx, email)
}

func createUserInAuth(ctx context.Context, fa *authentication.Auth, u *domain.User) (string, error) {
	return fa.CreateUser(ctx, u.Email, u.Password)
}

func setInFirestore(ctx context.Context, fs *firestore.Firestore, u *domain.User) error {
	return fs.Set(ctx, UserCollection, u.ID, u)
}
