package persistence

import (
	"context"
	"strings"

	"github.com/16francs/gran/api/user/middleware"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type userPersistence struct {
	auth      *authentication.Auth
	firestore *firestore.Firestore
}

const (
	// UserCollection - UserCollection名
	UserCollection = "users"
	// GroupCollection - GroupCollection名
	GroupCollection = "group"
)

// NewUserPersistence - UserRepositoryの生成
func NewUserPersistence(fa *authentication.Auth, fs *firestore.Firestore) repository.UserRepository {
	return &userPersistence{
		auth:      fa,
		firestore: fs,
	}
}

func (r *userPersistence) Authentication(ctx context.Context) (*domain.User, error) {
	t, err := getToken(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := r.auth.VerifyIDToken(ctx, t)
	if err != nil {
		return nil, err
	}

	doc, err := r.firestore.Client.Collection(UserCollection).Doc(uid).Get(ctx)
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

func (r *userPersistence) CreateGroup(ctx context.Context, u *domain.User) error {
	return nil
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

// TODO: リファクタ
func getUIDByEmailInAuth(ctx context.Context, fa *authentication.Auth, email string) (string, error) {
	return fa.GetUIDByEmail(ctx, email)
}

func createUserInAuth(ctx context.Context, fa *authentication.Auth, u *domain.User) (string, error) {
	return fa.CreateUser(ctx, u.Email, u.Password)
}

func setInFirestore(ctx context.Context, fs *firestore.Firestore, u *domain.User) error {
	return fs.Set(ctx, UserCollection, u.ID, u)
}
