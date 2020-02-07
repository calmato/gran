package request

// CreateBoard - ボード作成APIのリクエスト
type CreateBoard struct {
	Name            string   `json:"name" validate:"required"`
	GroupID         string   `json:"groupId" validate:"required"`
	Closed          bool     `json:"closed"`
	ThumbnailURL    string   `json:"thumbnailUrl"`
	BackgroundColor string   `json:"backgroundColor"`
	Labels          []string `json:"labels"`
}
