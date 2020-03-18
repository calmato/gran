package service

import (
	"context"
	"reflect"
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

func (brm *boardRepositoryMock) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       boardCurrent,
		UpdatedAt:       boardCurrent,
	}

	bs := []*domain.Board{b}

	return bs, nil
}

func (brm *boardRepositoryMock) Create(ctx context.Context, b *domain.Board) error {
	return nil
}

type fileUploaderMock struct{}

func (fum *fileUploaderMock) UploadBoardThumbnail(ctx context.Context, data []byte) (string, error) {
	return "test success", nil
}

func TestBoardService_Index(t *testing.T) {
	target := NewBoardService(&boardDomainValidationMock{}, &boardRepositoryMock{}, &fileUploaderMock{})

	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		IsClosed:        true,
		ThumbnailURL:    "",
		BackgroundColor: "",
		Labels:          make([]string, 0),
		GroupID:         "",
		CreatedAt:       boardCurrent,
		UpdatedAt:       boardCurrent,
	}

	want := []*domain.Board{b}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.Index(ctx, "JUA1ouY12ickxIupMVdVl3ieM7s2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestBoardService_Create(t *testing.T) {
	target := NewBoardService(&boardDomainValidationMock{}, &boardRepositoryMock{}, &fileUploaderMock{})

	b := &domain.Board{
		ID:              "JUA1ouY12ickxIupMVdVl3ieM7s2",
		Name:            "テストグループ",
		IsClosed:        true,
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

func TestBoardService_UploadThumbnail(t *testing.T) {
	target := NewBoardService(&boardDomainValidationMock{}, &boardRepositoryMock{}, &fileUploaderMock{})

	want := "test success"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	got, err := target.UploadThumbnail(ctx, []byte{})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
