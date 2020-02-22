package request

import (
	"mime/multipart"
)

// CreateBoard - ボード作成APIのリクエスト
type CreateBoard struct {
	Name            string         `json:"name" label:"ボード名" validate:"required,max=64"`
	GroupID         string         `json:"groupId" label:"グループID" validate:"required"`
	Closed          bool           `json:"closed" label:"非公開"`
	BackgroundColor string         `json:"backgroundColor" label:"背景色"`
	Labels          []string       `json:"labels" label:"ラベル一覧" validate:"unique,dive,max=16"`
	Thumbnail       multipart.File `label:"サムネイル"`
}
