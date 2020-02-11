package service

import (
	"context"
	"testing"
	"time"

	"github.com/16francs/gran/api/todo/internal/domain"
)

var boardCurrent = time.Now()

type boardDomainValidationMock struct{}

func (bdvm *boardDomainValidationMock) Board(ctx context.Context, g *domain.Board) []*domain.ValidationError {
	return nil
}

type boardRepositoryMock struct{}

func (brm *boardRepositoryMock) Create(ctx context.Context, b *domain.Board) error {
	return nil
}

func TestBoardService_Create(t *testing.T) {
	target := NewBoardService(&boardDomainValidationMock{}, &boardRepositoryMock{})

	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		Closed:          true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       boardCurrent,
		UpdatedAt:       boardCurrent,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := target.Create(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2", b)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
