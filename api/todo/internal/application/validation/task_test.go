package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestTaskRequestValidation_CreateTask(t *testing.T) {
	testCases := map[string]struct {
		Request  *request.CreateTask
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.CreateTask{
				Name:            "タスク",
				Description:     "説明",
				BoardListID:     "task-create-boardlist-id",
				Labels:          []string{},
				Attachments:     []string{},
				AssignedUserIDs: []string{},
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewTaskRequestValidation()

			got := target.CreateTask(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
