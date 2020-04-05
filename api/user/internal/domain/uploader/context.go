package uploader

import (
	"context"
)

// FileUploader - FilerUploaderインターフェース
type FileUploader interface {
	UploadUserThumbnail(ctx context.Context, data []byte) (string, error)
}
