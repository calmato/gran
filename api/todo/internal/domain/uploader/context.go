package uploader

import (
	"context"
	"mime/multipart"
)

// FileUploader - FilerUploaderインターフェース
type FileUploader interface {
	UploadBoardThumbnail(ctx context.Context, file multipart.File) (string, error)
}
