package storage

import (
	"context"

	"github.com/16francs/gran/api/user/internal/domain/uploader"
	gcs "github.com/16francs/gran/api/user/lib/firebase/storage"
)

type fileUploader struct {
	storage *gcs.Storage
}

// NewFileUploader - FileUploaderの生成
func NewFileUploader(cs *gcs.Storage) uploader.FileUploader {
	return &fileUploader{
		storage: cs,
	}
}

func (fu *fileUploader) UploadUserThumbnail(ctx context.Context, data []byte) (string, error) {
	thumbnailURL, err := fu.storage.Write(ctx, UserThumbnailPath, data)
	if err != nil {
		return "", err
	}

	return thumbnailURL, nil
}
