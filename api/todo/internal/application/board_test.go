package application

import (
	"context"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
)

var current = time.Now()

type boardRequestValidationMock struct{}

func (brvm *boardRequestValidationMock) CreateBoard(req *request.CreateBoard) []*domain.ValidationError {
	return nil
}

type boardServiceMock struct{}

func (bsm *boardServiceMock) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       current,
		UpdatedAt:       current,
	}

	bs := []*domain.Board{b}

	return bs, nil
}

func (bsm *boardServiceMock) Create(ctx context.Context, groupID string, b *domain.Board) error {
	return nil
}

type userServiceMock struct{}

func (usm *userServiceMock) Authentication(ctx context.Context) (*domain.User, error) {
	u := &domain.User{
		ID:           "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Email:        "hoge@hoge.com",
		Password:     "12345678",
		Name:         "テストユーザ",
		ThumbnailURL: "",
		GroupIDs:     make([]string, 0),
		CreatedAt:    current,
		UpdatedAt:    current,
	}

	return u, nil
}

func (usm *userServiceMock) IsContainInGroupIDs(ctx context.Context, groupID string, u *domain.User) bool {
	return true
}

func TestBoardApplication_Create(t *testing.T) {
	target := NewBoardApplication(&boardRequestValidationMock{}, &boardServiceMock{}, &userServiceMock{})

	b := &request.CreateBoard{
		Name:            "テストグループ",
		GroupID:         "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Create(ctx, b)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
