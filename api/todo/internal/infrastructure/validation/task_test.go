package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestTaskDomainValidation_Task(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Defined variable
	current := time.Now()

	task := &domain.Task{
		ID:              "task-id",
		Name:            "タスク",
		Description:     "説明",
		Labels:          []string{},
		AttachmentURLs:  []string{},
		BoardID:         "board-id",
		BoardListID:     "board-list-id",
		AssignedUserIDs: []string{},
		CheckListIDs:    []string{},
		CommentIDs:      []string{},
		DeadlinedAt:     current,
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	// Start test
	target := NewTaskDomainValidation()

	want := []*domain.ValidationError{}

	got := target.Task(ctx, task)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
