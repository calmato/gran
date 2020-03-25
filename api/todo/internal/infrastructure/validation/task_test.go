package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestTaskDomainValidation_Task(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Task     *domain.Task
		Expected []*domain.ValidationError
	}{
		"ok": {
			Task: &domain.Task{
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
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskDomainValidation()

			got := target.Task(ctx, testCase.Task)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
