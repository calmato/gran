package firestore

import (
	"context"

	"golang.org/x/xerrors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// Firestore - Firestoreの構造体
type Firestore struct {
	Client *firestore.Client
}

// NewClient - Firestoreに接続
func NewClient(ctx context.Context, app *firebase.App) (*Firestore, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Failed to NewFirestore: %w", err)
	}

	return &Firestore{client}, nil
}

// Close - Firestoreとの接続を終了
func (f *Firestore) Close() error {
	return f.Client.Close()
}
