package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
)

func TestGroupDomainValidation_Group(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Defined variable
	current := time.Now()

	// Defined variable
	g := &domain.Group{
		ID:          "group-id",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	// Start test
	target := NewGroupDomainValidation()

	want := []*domain.ValidationError{}

	got := target.Group(ctx, g)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
