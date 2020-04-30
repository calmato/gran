package persistence

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/domain/repository"
	"github.com/calmato/gran/api/todo/lib/firebase/firestore"
)

type taskPersistence struct {
	firestore *firestore.Firestore
}

// NewTaskPersistence - TaskRepositoryの生成
func NewTaskPersistence(fs *firestore.Firestore) repository.TaskRepository {
	return &taskPersistence{
		firestore: fs,
	}
}

func (tp *taskPersistence) IndexByBoardID(ctx context.Context, boardID string) ([]*domain.Task, error) {
	ts := make([]*domain.Task, 0)
	taskCollection := getTaskCollection()

	q := &firestore.Query{
		Field:    "board_id", // TODO: Taskエンティティのタグから取得
		Operator: "==",
		Value:    boardID,
	}

	iter := tp.firestore.GetByQuery(ctx, taskCollection, q)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		t := &domain.Task{}

		err = doc.DataTo(t)
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func (tp *taskPersistence) Show(ctx context.Context, taskID string) (*domain.Task, error) {
	taskCollection := getTaskCollection()

	doc, err := tp.firestore.Get(ctx, taskCollection, taskID)
	if err != nil {
		return nil, err
	}

	t := &domain.Task{}

	err = doc.DataTo(t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (tp *taskPersistence) Create(ctx context.Context, t *domain.Task) error {
	taskCollection := getTaskCollection()

	if err := tp.firestore.Set(ctx, taskCollection, t.ID, t); err != nil {
		return err
	}

	return nil
}
