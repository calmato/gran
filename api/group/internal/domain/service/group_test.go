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
	current := time.Now()

	testCases := map[string]struct {
		User     *domain.User
		Expected []*domain.Group
	}{
		"ok": {
			User: &domain.User{
				ID:        "user-id",
				Email:     "hoge@hoge.com",
				GroupIDs:  []string{"group-id"},
				CreatedAt: current,
				UpdatedAt: current,
			},
			Expected: []*domain.Group{
				{
					ID:          "group-id",
					Name:        "テストグループ",
					Description: "グループの説明",
					UserIDs:     []string{"user-id"},
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

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Index(ctx, testCase.User).Return(testCase.Expected, nil)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			got, err := target.Index(ctx, testCase.User)
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

func TestGroupService_Show(t *testing.T) {
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

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Show(ctx, testCase.GroupID).Return(testCase.Expected, nil)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			got, err := target.Show(ctx, testCase.GroupID)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestGroupService_Create(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		User  *domain.User
		Group *domain.Group
	}{
		"ok": {
			User: &domain.User{
				ID:        "user-id",
				Email:     "hoge@hoge.com",
				GroupIDs:  []string{"group-id"},
				CreatedAt: current,
				UpdatedAt: current,
			},
			Group: &domain.Group{
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
		ves := make([]*domain.ValidationError, 0)

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
		gdvm.EXPECT().Group(ctx, testCase.Group).Return(ves)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Create(ctx, testCase.Group).Return(nil)

		urm := mock_repository.NewMockUserRepository(ctrl)
		urm.EXPECT().Update(ctx, testCase.User).Return(nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			got, err := target.Create(ctx, testCase.User, testCase.Group)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, testCase.Group) {
				t.Fatalf("want %#v, but %#v", testCase.Group, got)
				return
			}
		})
	}
}

func TestGroupService_Update(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Group *domain.Group
	}{
		"ok": {
			Group: &domain.Group{
				ID:          "group-id",
				Name:        "テストグループ",
				Description: "グループの説明",
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
		ves := make([]*domain.ValidationError, 0)

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
		gdvm.EXPECT().Group(ctx, testCase.Group).Return(ves)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Update(ctx, testCase.Group).Return(nil)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			err := target.Update(ctx, testCase.Group)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupService_InviteUsers(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Group *domain.Group
	}{
		"ok": {
			Group: &domain.Group{
				ID:            "group-id",
				Name:          "テストグループ",
				Description:   "グループの説明",
				InvitedEmails: []string{"user01@hoge.com"},
				BoardIDs:      make([]string, 0),
				CreatedAt:     current,
				UpdatedAt:     current,
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

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
		gdvm.EXPECT().Group(ctx, testCase.Group).Return(ves)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Update(ctx, testCase.Group).Return(nil)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			err := target.InviteUsers(ctx, testCase.Group)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupService_Join(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Group *domain.Group
	}{
		"ok": {
			Group: &domain.Group{
				ID:            "group-id",
				Name:          "テストグループ",
				Description:   "グループの説明",
				UserIDs:       []string{"user-id"},
				InvitedEmails: []string{},
				BoardIDs:      make([]string, 0),
				CreatedAt:     current,
				UpdatedAt:     current,
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

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)
		gdvm.EXPECT().Group(ctx, testCase.Group).Return(ves)

		grm := mock_repository.NewMockGroupRepository(ctrl)
		grm.EXPECT().Update(ctx, testCase.Group).Return(nil)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			err := target.Join(ctx, testCase.Group)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestGroupService_IsContainInUserIDs(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UserID string
		Group  *domain.Group
	}{
		"ok": {
			UserID: "user-id",
			Group: &domain.Group{
				ID:            "group-id",
				Name:          "テストグループ",
				Description:   "グループの説明",
				UserIDs:       []string{"user-id"},
				InvitedEmails: make([]string, 0),
				BoardIDs:      make([]string, 0),
				CreatedAt:     current,
				UpdatedAt:     current,
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

		grm := mock_repository.NewMockGroupRepository(ctrl)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			got := target.IsContainInUserIDs(ctx, testCase.UserID, testCase.Group)
			if !got {
				t.Fatalf("error: %v", got)
				return
			}
		})
	}
}

func TestGroupService_IsContainInInvitedEmails(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Email string
		Group *domain.Group
	}{
		"ok": {
			Email: "hoge@hoge.com",
			Group: &domain.Group{
				ID:            "group-id",
				Name:          "テストグループ",
				Description:   "グループの説明",
				UserIDs:       make([]string, 0),
				InvitedEmails: []string{"hoge@hoge.com"},
				BoardIDs:      make([]string, 0),
				CreatedAt:     current,
				UpdatedAt:     current,
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined mocks
		gdvm := mock_validation.NewMockGroupDomainValidation(ctrl)

		grm := mock_repository.NewMockGroupRepository(ctrl)

		urm := mock_repository.NewMockUserRepository(ctrl)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupService(gdvm, grm, urm)

			got := target.IsContainInInvitedEmails(ctx, testCase.Email, testCase.Group)
			if !got {
				t.Fatalf("error: %v", got)
				return
			}
		})
	}
}
