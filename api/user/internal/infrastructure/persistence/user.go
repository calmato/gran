package persistence

import (
	"context"
	"strings"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
	"github.com/16francs/gran/api/user/middleware"
)

type userPersistence struct {
	auth      *authentication.Auth
	firestore *firestore.Firestore
}

// NewUserPersistence - UserRepositoryの生成
func NewUserPersistence(fa *authentication.Auth, fs *firestore.Firestore) repository.UserRepository {
	return &userPersistence{
		auth:      fa,
		firestore: fs,
	}
}

func (up *userPersistence) Authentication(ctx context.Context) (*domain.User, error) {
	t, err := getToken(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := up.auth.VerifyIDToken(ctx, t)
	if err != nil {
		return nil, err
	}

	doc, err := up.firestore.Client.Collection(UserCollection).Doc(uid).Get(ctx)
	if err != nil {
		return nil, err
	}

	u := &domain.User{}

	// TODO: メソッド化
	err = doc.DataTo(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (up *userPersistence) Create(ctx context.Context, u *domain.User) error {
	uid, err := up.auth.CreateUser(ctx, u.Email, u.Password)
	if err != nil {
		return err
	}

	u.ID = uid

	if err = up.firestore.Set(ctx, UserCollection, u.ID, u); err != nil {
		return err
	}

	return nil
}

func (up *userPersistence) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := up.auth.GetUIDByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func getToken(ctx context.Context) (string, error) {
	gc, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return "", xerrors.New("Cannot convert to gin.Context")
	}

	a := gc.GetHeader("Authorization")
	if a == "" {
		return "", xerrors.New("Authorization Header is not contain.")
	}

	t := strings.Replace(a, "Bearer ", "", 1)
	return t, nil
}
