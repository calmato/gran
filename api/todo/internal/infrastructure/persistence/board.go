package persistence

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/repository"
	"github.com/calmato/gran/api/todo/lib/firebase/firestore"
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
	boardCollection := getBoardCollection(groupID)

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

func (bp *boardPersistence) Show(ctx context.Context, groupID string, boardID string) (*domain.Board, error) {
	boardCollection := getBoardCollection(groupID)

	doc, err := bp.firestore.Get(ctx, boardCollection, boardID)
	if err != nil {
		return nil, err
	}

	b := &domain.Board{}

	err = doc.DataTo(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bp *boardPersistence) Create(ctx context.Context, b *domain.Board) error {
	boardCollection := getBoardCollection(b.GroupID)

	if err := bp.firestore.Set(ctx, boardCollection, b.ID, b); err != nil {
		return err
	}

	return nil
}

func (bp *boardPersistence) Update(ctx context.Context, b *domain.Board) error {
	boardCollection := getBoardCollection(b.GroupID)

	if err := bp.firestore.Set(ctx, boardCollection, b.ID, b); err != nil {
		return err
	}

	return nil
}

func (bp *boardPersistence) IndexBoardList(
	ctx context.Context, groupID string, boardID string,
) ([]*domain.BoardList, error) {
	bls := make([]*domain.BoardList, 0)
	boardListCollection := getBoardListCollection(groupID, boardID)

	iter := bp.firestore.GetAll(ctx, boardListCollection)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		bl := &domain.BoardList{}

		err = doc.DataTo(bl)
		if err != nil {
			return nil, err
		}

		bls = append(bls, bl)
	}

	return bls, nil
}

func (bp *boardPersistence) ShowBoardList(
	ctx context.Context, groupID string, boardID string, boardListID string,
) (*domain.BoardList, error) {
	boardListCollection := getBoardListCollection(groupID, boardID)

	doc, err := bp.firestore.Get(ctx, boardListCollection, boardListID)
	if err != nil {
		return nil, err
	}

	bl := &domain.BoardList{}

	err = doc.DataTo(bl)
	if err != nil {
		return nil, err
	}

	return bl, nil
}

func (bp *boardPersistence) CreateBoardList(
	ctx context.Context, groupID string, boardID string, bl *domain.BoardList,
) error {
	boardListCollection := getBoardListCollection(groupID, boardID)

	if err := bp.firestore.Set(ctx, boardListCollection, bl.ID, bl); err != nil {
		return err
	}

	return nil
}

func (bp *boardPersistence) UpdateBoardList(
	ctx context.Context, groupID string, boardID string, bl *domain.BoardList,
) error {
	boardListCollection := getBoardListCollection(groupID, boardID)

	if err := bp.firestore.Set(ctx, boardListCollection, bl.ID, bl); err != nil {
		return err
	}

	return nil
}
