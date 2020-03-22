package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
)

var current = time.Now()

func TestBoardDomainValidation_Board(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Defined variable
	b := &domain.Board{
		ID:              "board-id",
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	// Start test
	target := NewBoardDomainValidation()

	want := []*domain.ValidationError{}

	got := target.Board(ctx, b)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardDomainValidation_BoardList(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Defined variables
	bl := &domain.BoardList{
		ID:        "boardlist-id",
		Name:      "テストボードリスト",
		BoardID:   "",
		CreatedAt: current,
		UpdatedAt: current,
	}

	// Start test
	target := NewBoardDomainValidation()

	want := []*domain.ValidationError{}

	got := target.BoardList(ctx, bl)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
