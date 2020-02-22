package uploader

import (
	"context"
)

// FileUploader - FilerUploaderインターフェース
type FileUploader interface {
	UploadBoardThumbnail(ctx context.Context, data []byte) (string, error)
}
