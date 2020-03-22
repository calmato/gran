package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestBoardRequestValidation_CreateBoard(t *testing.T) {
	// Defined variables
	b := &request.CreateBoard{
		Name:            "テストグループ",
		IsClosed:        true,
		Thumbnail:       "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
	}

	// Start test
	target := NewBoardRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.CreateBoard(b)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardRequestValidation_CreateBoardList(t *testing.T) {
	// Defined variables
	bl := &request.CreateBoardList{
		Name:  "テストボードグループ",
		Color: "",
	}

	// Start test
	target := NewBoardRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.CreateBoardList(bl)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
