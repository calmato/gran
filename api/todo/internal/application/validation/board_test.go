package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

func TestBoardRequestValidation_CreateBoard(t *testing.T) {
	target := NewBoardRequestValidation()

	want := []*domain.ValidationError(nil)

	b := &request.CreateBoard{
		Name:            "テストグループ",
		GroupID:         "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Closed:          true,
		Thumbnail:       "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
	}

	got := target.CreateBoard(b)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
