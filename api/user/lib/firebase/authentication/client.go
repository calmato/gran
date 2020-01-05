package authentication

import (
	"context"

	"golang.org/x/xerrors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// Authentication - Authenticationの構造体
type Authentication struct {
	Client *auth.Client
}

// NewClient - Firebase Authenticationに接続
func NewClient(ctx context.Context, app *firebase.App) (*Authentication, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Failed to NewAuthentication: %w", err)
	}

	return &Authentication{client}, nil
}

// VerifyIDToken - IDトークンを確認して、デコードされたトークンからデバイスのuidを取得
func (a *Authentication) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	t, err := a.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return t.UID, nil
}

// GetUIDByEmail - メールアドレスによるユーザーUIDの取得
func (a *Authentication) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	u, err := a.Client.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// CreateUser - 新しいFirebase Authenticationユーザーを作成
func (a *Authentication) CreateUser(ctx context.Context, email string, password string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(emailVerified(email)).
		Password(password).
		Disabled(false)

	u, err := a.Client.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// UpdateUser - 既存のユーザーのデータを変更
func (a *Authentication) UpdateUser(ctx context.Context, uid string, email string, password string, disable bool) error {
	params := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(emailVerified(email)).
		Password(password).
		Disabled(disable)

	if _, err := a.Client.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

// DeleteUser - 既存のユーザーをuidで削除
func (a *Authentication) DeleteUser(ctx context.Context, uid string) error {
	return a.Client.DeleteUser(ctx, uid)
}

func emailVerified(email string) bool {
	return email != ""
}
