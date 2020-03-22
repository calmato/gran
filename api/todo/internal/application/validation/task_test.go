package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestTaskRequestValidation_CreateTask(t *testing.T) {
	// Defined variables
	task := &request.CreateTask{
		Name:            "タスク",
		Description:     "説明",
		BoardListID:     "task-create-boardlist-id",
		Labels:          []string{},
		Attachments:     []string{},
		AssignedUserIDs: []string{},
	}

	// Start test
	target := NewTaskRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.CreateTask(task)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
