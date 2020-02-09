package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
)

var current = time.Now()

type groupRepositoryMock struct{}

func (grm *groupRepositoryMock) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserRefs:    make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	gs := []*domain.Group{g}

	return gs, nil
}

func (grm *groupRepositoryMock) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserRefs:    make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	return g, nil
}

func (grm *groupRepositoryMock) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	return nil
}

func (grm *groupRepositoryMock) Update(ctx context.Context, g *domain.Group) error {
	return nil
}

func (grm *groupRepositoryMock) UserIDExistsInUserRefs(ctx context.Context, userID string, g *domain.Group) bool {
	return true
}

func TestGroupDomainValidation_Group(t *testing.T) {
	target := NewGroupDomainValidation(&groupRepositoryMock{})

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
