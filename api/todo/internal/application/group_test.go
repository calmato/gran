package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
	mock_validation "github.com/16francs/gran/api/todo/mock/application/validation"
	mock_service "github.com/16francs/gran/api/todo/mock/domain/service"
)

func TestGroupApplication_Index(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Expected []*domain.Group
	}{
		"ok": {
			Expected: []*domain.Group{
				{
					ID:          "group-id",
					Name:        "テストグループ",
					Description: "グループの説明",
					UserIDs:     make([]string, 0),
					BoardIDs:    make([]string, 0),
					CreatedAt:   current,
					UpdatedAt:   current,
				},
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		u := &domain.User{}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Index(ctx, u).Return(testCase.Expected, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			got, err := target.Index(ctx)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestGroupApplication_Show(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID  string
		Expected *domain.Group
	}{
		"ok": {
			GroupID: "group-id",
			Expected: &domain.Group{
				ID:          "group-id",
				Name:        "テストグループ",
				Description: "グループの説明",
				UserIDs:     []string{"user-id"},
				BoardIDs:    make([]string, 0),
				CreatedAt:   current,
				UpdatedAt:   current,
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		u := &domain.User{
			ID:       "user-id",
			GroupIDs: []string{"group-id"},
		}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Show(ctx, testCase.GroupID).Return(testCase.Expected, nil)
		gsm.EXPECT().IsContainInUserIDs(ctx, u.ID, testCase.Expected).Return(true)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			got, err := target.Show(ctx, testCase.GroupID)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
			}
		})
	}
}

func TestGroupApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Request *request.CreateGroup
	}{
		"ok": {
			Request: &request.CreateGroup{
				Name:        "テストグループ",
				Description: "説明",
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{}

		g := &domain.Group{
			Name:        testCase.Request.Name,
			Description: testCase.Request.Description,
		}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
		grvm.EXPECT().CreateGroup(testCase.Request).Return(ves)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Create(ctx, u, g).Return(g, nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			err := target.Create(ctx, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupApplication_Update(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID string
		Request *request.UpdateGroup
	}{
		"ok": {
			GroupID: "group-id",
			Request: &request.UpdateGroup{
				Name:        "テストグループ",
				Description: "説明",
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{
			ID:       "user-id",
			GroupIDs: []string{"group-id"},
		}

		g := &domain.Group{
			ID:          testCase.GroupID,
			Name:        testCase.Request.Name,
			Description: testCase.Request.Description,
			UserIDs:     []string{"user-id"},
			BoardIDs:    make([]string, 0),
			CreatedAt:   current,
			UpdatedAt:   current,
		}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
		grvm.EXPECT().UpdateGroup(testCase.Request).Return(ves)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Show(ctx, testCase.GroupID).Return(g, nil)
		gsm.EXPECT().Update(ctx, g).Return(nil)
		gsm.EXPECT().IsContainInUserIDs(ctx, u.ID, g).Return(true)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			err := target.Update(ctx, testCase.GroupID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupApplication_InviteUsers(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID string
		Request *request.InviteUsers
	}{
		"ok": {
			GroupID: "group-id",
			Request: &request.InviteUsers{
				Emails: []string{"hoge@hoge.com"},
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{
			ID:       "user-id",
			GroupIDs: []string{"group-id"},
		}

		g := &domain.Group{
			ID:            testCase.GroupID,
			Name:          "テストグループ",
			Description:   "説明",
			UserIDs:       []string{"user-id"},
			InvitedEmails: make([]string, 0),
			BoardIDs:      make([]string, 0),
			CreatedAt:     current,
			UpdatedAt:     current,
		}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)
		grvm.EXPECT().InviteUsers(testCase.Request).Return(ves)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Show(ctx, testCase.GroupID).Return(g, nil)
		gsm.EXPECT().InviteUsers(ctx, g).Return(nil)
		gsm.EXPECT().IsContainInUserIDs(ctx, u.ID, g).Return(true)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			err := target.InviteUsers(ctx, testCase.GroupID, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupApplication_Join(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		GroupID string
	}{
		"ok": {
			GroupID: "group-id",
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		u := &domain.User{
			ID:       "user-id",
			Email:    "hoge@hoge.com",
			GroupIDs: []string{"group-id"},
		}

		g := &domain.Group{
			ID:            testCase.GroupID,
			Name:          "テストグループ",
			Description:   "説明",
			UserIDs:       []string{"user-id"},
			InvitedEmails: []string{"hoge@hoge.com"},
			BoardIDs:      make([]string, 0),
			CreatedAt:     current,
			UpdatedAt:     current,
		}

		// Defined mocks
		grvm := mock_validation.NewMockGroupRequestValidation(ctrl)

		gsm := mock_service.NewMockGroupService(ctrl)
		gsm.EXPECT().Show(ctx, testCase.GroupID).Return(g, nil)
		gsm.EXPECT().IsContainInInvitedEmails(ctx, u.Email, g).Return(true)
		gsm.EXPECT().Join(ctx, g).Return(nil)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupApplication(grvm, gsm, usm)

			err := target.Join(ctx, testCase.GroupID)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
