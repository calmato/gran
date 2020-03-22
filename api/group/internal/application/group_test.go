package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
	mock_validation "github.com/16francs/gran/api/group/mock/application/validation"
	mock_service "github.com/16francs/gran/api/group/mock/domain/service"
)

func TestGroupApplication_Index(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()

	u := &domain.User{}

	g := &domain.Group{
		ID:          "group-index-group-id",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	gs := []*domain.Group{g}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Index(ctx, u).Return(gs, nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	want := gs

	got, err := target.Index(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupApplication_Show(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-show-group-id"
	userID := "board-show-user-id"

	u := &domain.User{
		ID: userID,
	}

	g := &domain.Group{
		ID:          "group-show-group-id",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     []string{userID},
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Show(ctx, groupID).Return(g, nil)
	gsm.EXPECT().IsContainInUserIDs(ctx, userID, g).Return(true)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	want := g

	got, err := target.Show(ctx, groupID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupApplication_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	ves := make([]*domain.ValidationError, 0)

	req := &request.CreateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	u := &domain.User{}

	g := &domain.Group{
		Name:        "テストグループ",
		Description: "説明",
	}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
	grvm.EXPECT().CreateGroup(req).Return(ves)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Create(ctx, u, g).Return(nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	err := target.Create(ctx, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupApplication_Update(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-update-group-id"
	userID := "board-update-user-id"
	ves := make([]*domain.ValidationError, 0)

	req := &request.UpdateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	u := &domain.User{
		ID: userID,
	}

	g := &domain.Group{
		ID:          "group-index-group-id",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
	grvm.EXPECT().UpdateGroup(req).Return(ves)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Show(ctx, groupID).Return(g, nil)
	gsm.EXPECT().Update(ctx, g).Return(nil)
	gsm.EXPECT().IsContainInUserIDs(ctx, userID, g).Return(true)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	err := target.Update(ctx, groupID, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupApplication_InviteUsers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-inviteusers-group-id"
	userID := "board-inviteusers-user-id"
	email := "hoge@hoge.com"
	ves := make([]*domain.ValidationError, 0)

	req := &request.InviteUsers{
		Emails: []string{email},
	}

	u := &domain.User{
		ID: userID,
	}

	g := &domain.Group{
		ID:          "group-index-group-id",
		Name:        "テストグループ",
		Description: "グループの説明",
		UserIDs:     make([]string, 0),
		CreatedAt:   current,
		UpdatedAt:   current,
	}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
	grvm.EXPECT().InviteUsers(req).Return(ves)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Show(ctx, groupID).Return(g, nil)
	gsm.EXPECT().InviteUsers(ctx, g).Return(nil)
	gsm.EXPECT().IsContainInUserIDs(ctx, userID, g).Return(true)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	err := target.InviteUsers(ctx, groupID, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestGroupApplication_Join(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defined variables
	current := time.Now()
	groupID := "board-join-group-id"
	userID := "board-join-user-id"
	email := "hoge@hoge.com"

	u := &domain.User{
		ID:    userID,
		Email: email,
	}

	g := &domain.Group{
		ID:            "group-join-group-id",
		Name:          "テストグループ",
		Description:   "グループの説明",
		UserIDs:       []string{},
		InvitedEmails: []string{email},
		CreatedAt:     current,
		UpdatedAt:     current,
	}

	// Defined mocks
	grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

	gsm := mock_service.NewMockGroupService(ctrl)
	gsm.EXPECT().Show(ctx, groupID).Return(g, nil)
	gsm.EXPECT().IsContainInInvitedEmails(ctx, email, g).Return(true)
	gsm.EXPECT().Join(ctx, g).Return(nil)

	usm := mock_service.NewMockUserService(ctrl)
	usm.EXPECT().Authentication(ctx).Return(u, nil)

	// Start test
	target := NewGroupApplication(grvm, gsm, usm)

	err := target.Join(ctx, groupID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
