package validation

import (
	"reflect"
	"testing"

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
