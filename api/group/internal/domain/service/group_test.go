package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/16francs/gran/api/group/internal/domain"
	mock_repository "github.com/16francs/gran/api/group/mock/domain/repository"
	mock_validation "github.com/16francs/gran/api/group/mock/domain/validation"
	"github.com/golang/mock/gomock"
)

func TestGroupService_Index(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "group-show-group-id"
	userID := "group-show-user-id"

	u := &domain.User{
		ID:       userID,
		GroupIDs: []string{groupID},
	}

	g := &domain.Group{
		ID:          groupID,
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     []string{userID},
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	gs := []*domain.Group{g}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Index(ctx, u).Return(gs, nil)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	want := gs

	got, err := target.Index(ctx, u)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupService_Show(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "group-show-group-id"

	g := &domain.Group{
		ID:          groupID,
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Show(ctx, groupID).Return(g, nil)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	want := g

	got, err := target.Show(ctx, groupID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupService_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
	}

	u := &domain.User{}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
	gdvm.EXPECT().Group(ctx, g).Return(ves)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Create(ctx, g).Return(nil)

	urm := mock_repository.NewMockUserRepository(ctrl)
	urm.EXPECT().Update(ctx, u).Return(nil)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	err := target.Create(ctx, u, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_Update(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
	gdvm.EXPECT().Group(ctx, g).Return(ves)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Update(ctx, g).Return(nil)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	err := target.Update(ctx, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_InviteUsers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
	gdvm.EXPECT().Group(ctx, g).Return(ves)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Update(ctx, g).Return(nil)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	err := target.InviteUsers(ctx, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_Join(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
	gdvm.EXPECT().Group(ctx, g).Return(ves)

	grm := mock_repository.NewMockGroupRepository(ctrl)
	grm.EXPECT().Update(ctx, g).Return(nil)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	err := target.Join(ctx, g)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupService_IsContainInUserIDs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	userID := "group-iscontain-user-id"

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     []string{userID},
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

	grm := mock_repository.NewMockGroupRepository(ctrl)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	got := target.IsContainInUserIDs(ctx, userID, g)
	if !got {
		t.Fatalf("error: %v", got)
	}
}

func TestGroupService_IsContainInInvitedEmails(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	email := "hoge@hoge.com"

	g := &domain.Group{
		Name:          "テストグループ",
		Description:   "グループの説明",
		InvitedEmails: []string{email},
	}

	// Defined mocks
	gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

	grm := mock_repository.NewMockGroupRepository(ctrl)

	urm := mock_repository.NewMockUserRepository(ctrl)

	// Start test
	target := NewGroupService(gdvm, grm, urm)

	got := target.IsContainInInvitedEmails(ctx, email, g)
	if !got {
		t.Fatalf("error: %v", got)
	}
}
