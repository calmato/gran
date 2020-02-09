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
	target := NewBoardDomainValidation()

	want := []*domain.ValidationError{}

	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupRef:        "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got := target.Board(ctx, b)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
