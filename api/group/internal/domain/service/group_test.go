package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
)

var (
	groupCurrent  = time.Now()
	groupAuthUser = &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    groupCurrent,
		UpdatedAt:    groupCurrent,
	}
)

type groupDomainValidationMock struct{}

func (gdvm *groupDomainValidationMock) Group(ctx context.Context, g *domain.Group) []*domain.ValidationError {
	return nil
}

type groupRepositoryMock struct{}

func (grm *groupRepositoryMock) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	gs := []*domain.Group{g}

	return gs, nil
}

func (grm *groupRepositoryMock) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	return g, nil
}

func (grm *groupRepositoryMock) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	return nil
}

func (grm *groupRepositoryMock) Update(ctx context.Context, g *domain.Group) error {
	return nil
}

func TestGroupService_Index(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	want := []*domain.Group{g}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.Index(ctx, groupAuthUser)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupService_Show(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	want := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.Show(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupService_Create(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Create(ctx, groupAuthUser, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_Update(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Update(ctx, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_IsContainInUserIDs(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got := target.IsContainInUserIDs(ctx, groupAuthUser.ID, g)
	if !got {
		t.Fatalf("error: %v", got)
	}
}

func TestGroupService_IsContainInInvitedEmails(t *testing.T) {
	target := NewGroupService(&groupDomainValidationMock{}, &groupRepositoryMock{})

	email := "hoge@hoge.com"

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   groupCurrent,
		UpdatedAt:   groupCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got := target.IsContainInInvitedEmails(ctx, email, g)
	if !got {
		t.Fatalf("error: %v", got)
	}
}
