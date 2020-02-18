package persistence

import (
	"context"

	"google.golang.org/api/iterator"

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

func (bp *boardPersistence) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	bs := make([]*domain.Board, 0)
	boardCollection := GetBoardCollection(groupID)

	iter := bp.firestore.GetAll(ctx, boardCollection)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		b := &domain.Board{}

		err = doc.DataTo(b)
		if err != nil {
			return nil, err
		}

		bs = append(bs, b)
	}

	return bs, nil
}

func (bp *boardPersistence) Create(ctx context.Context, b *domain.Board) error {
	boardCollection := GetBoardCollection(b.GroupID)

	if err := bp.firestore.Set(ctx, boardCollection, b.ID, b); err != nil {
		return err
	}

	return nil
}
