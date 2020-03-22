package request

// CreateBoard - ボード作成APIのリクエスト
type CreateBoard struct {
	Name            string   `json:"name" label:"ボード名" validate:"required,max=64"`
	IsClosed        bool     `json:"isClosed" label:"非公開"`
	BackgroundColor string   `json:"backgroundColor" label:"背景色"`
	Thumbnail       string   `json:"thumbnail" label:"サムネイル"`
	Labels          []string `json:"labels" label:"ラベル一覧" validate:"unique,dive,max=16"`
}
