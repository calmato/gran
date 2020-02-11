package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
)

var current = time.Now()

func TestGroupDomainValidation_Group(t *testing.T) {
	target := NewGroupDomainValidation()

	want := []*domain.ValidationError{}

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserRefs:    make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got := target.Group(ctx, g)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
