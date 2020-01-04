package persistence

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type userRepository struct {
	*firestore.Firestore
}

// UserCollection - UserCollection名
const UserCollection = "users"

// NewUserRepository - UserRepositoryの生成
func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	// TODO: addInFirebaseの実装

	if err := addInFirestore(ctx, r.Firestore, u); err != nil {
		return err
	}

	return nil
}

func addInFirestore(ctx context.Context, fs *firestore.Firestore, u *domain.User) error {
	return fs.Add(ctx, UserCollection, u)
}
