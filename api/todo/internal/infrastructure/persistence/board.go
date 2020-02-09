package persistence

import (
	"context"

	"github.com/google/uuid"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
)

type boardPersistence struct {
	firestore *firestore.Firestore
}

// NewBoardPersistence - BoardRepositoryの生成
func NewBoardPersistence(fs *firestore.Firestore) repository.BoardRepository {
	return &boardPersistence{
		firestore: fs,
	}
}

func (bp *boardPersistence) Create(ctx context.Context, groupID string, b *domain.Board) error {
	b.ID = uuid.New().String()
	b.GroupRef = getGroupReference(groupID)

	if err := bp.firestore.Set(ctx, getBoardCollection(b.GroupRef), b.ID, b); err != nil {
		return err
	}

	return nil
}
