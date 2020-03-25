package validation

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestBoardRequestValidation_CreateBoard(t *testing.T) {
	testCases := map[string]struct {
		Request  *request.CreateBoard
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.CreateBoard{
				Name:            "テストグループ",
				IsClosed:        true,
				Thumbnail:       "",
				BackgroundColor: "",
				Labels:          make([]string, 0),
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardRequestValidation()

			got := target.CreateBoard(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardRequestValidation_CreateBoardList(t *testing.T) {
	testCases := map[string]struct {
		Request  *request.CreateBoardList
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.CreateBoardList{
				Name:  "テストボードグループ",
				Color: "",
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardRequestValidation()

			got := target.CreateBoardList(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardRequestValidation_UpdateKanban(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Request  *request.UpdateKanban
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.UpdateKanban{
				Lists: []*request.ListInUpdateKanban{
					{
						ID:    "board-list-id-01",
						Name:  "テストボードリスト01",
						Color: "red",
						Tasks: []*request.TaskInUpdateKanban{
							{
								ID:              "task-id",
								Name:            "テストタスク",
								Labels:          make([]string, 0),
								AssignedUserIDs: make([]string, 0),
								DeadlinedAt:     current,
							},
						},
					},
					{
						ID:    "board-list-id-02",
						Name:  "テストボードリスト02",
						Color: "red",
						Tasks: make([]*request.TaskInUpdateKanban, 0),
					},
				},
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardRequestValidation()

			got := target.UpdateKanban(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
