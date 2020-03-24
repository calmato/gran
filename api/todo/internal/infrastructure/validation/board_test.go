package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestBoardDomainValidation_Board(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Board    *domain.Board
		Expected []*domain.ValidationError
	}{
		"ok": {
			Board: &domain.Board{
				ID:              "board-id",
				Name:            "テストグループ",
				IsClosed:        true,
				ThumbnailURL:    "",
				BackgroundColor: "",
				Labels:          make([]string, 0),
				GroupID:         "",
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
			target := NewBoardDomainValidation()

			got := target.Board(ctx, testCase.Board)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestBoardDomainValidation_BoardList(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		BoardList *domain.BoardList
		Expected  []*domain.ValidationError
	}{
		"ok": {
			BoardList: &domain.BoardList{
				ID:        "board-list-id",
				Name:      "テストボードリスト",
				BoardID:   "",
				CreatedAt: current,
				UpdatedAt: current,
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewBoardDomainValidation()

			got := target.BoardList(ctx, testCase.BoardList)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
