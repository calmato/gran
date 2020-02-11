package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
)

var current = time.Now()

type groupRequestValidationMock struct{}

func (grvm *groupRequestValidationMock) CreateGroup(req *request.CreateGroup) []*domain.ValidationError {
	return nil
}

func (grvm *groupRequestValidationMock) UpdateGroup(req *request.UpdateGroup) []*domain.ValidationError {
	return nil
}

func (grvm *groupRequestValidationMock) InviteUsers(req *request.InviteUsers) []*domain.ValidationError {
	return nil
}

type groupServiceMock struct{}

func (gsm *groupServiceMock) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
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

func (gsm *groupServiceMock) Show(ctx context.Context, groupID string) (*domain.Group, error) {
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

func (gsm *groupServiceMock) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	return nil
}

func (gsm *groupServiceMock) Update(ctx context.Context, g *domain.Group) error {
	return nil
}

func (gsm *groupServiceMock) InviteUsers(ctx context.Context, g *domain.Group) error {
	return nil
}

func (gsm *groupServiceMock) UserIDExistsInUserRefs(ctx context.Context, userID string, g *domain.Group) bool {
	return true
}

func (gsm *groupServiceMock) EmailExistsInInvitedEmails(ctx context.Context, email string, g *domain.Group) bool {
	return true
}

type userServiceMock struct{}

func (usm *userServiceMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupRefs:    make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	return u, nil
}

func TestGroupApplication_Index(t *testing.T) {
	target := NewGroupApplication(&groupRequestValidationMock{}, &groupServiceMock{}, &userServiceMock{})

	g := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserRefs:    make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	want := []*domain.Group{g}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.Index(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupApplication_Show(t *testing.T) {
	target := NewGroupApplication(&groupRequestValidationMock{}, &groupServiceMock{}, &userServiceMock{})

	want := &domain.Group{
		ID:          "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserRefs:    make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
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

func TestGroupApplication_Create(t *testing.T) {
	target := NewGroupApplication(&groupRequestValidationMock{}, &groupServiceMock{}, &userServiceMock{})

	g := &request.CreateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Create(ctx, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupApplication_Update(t *testing.T) {
	target := NewGroupApplication(&groupRequestValidationMock{}, &groupServiceMock{}, &userServiceMock{})

	g := &request.UpdateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Update(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s", g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
