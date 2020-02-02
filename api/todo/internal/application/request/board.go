package request

// CreateBoard - ボード作成APIのリクエスト
type CreateBoard struct {
	Name            string   `json:"name" validate:"required"`
	GroupID         string   `json:"groupId" validate:"required"`
	Closed          bool     `firestore:"closed" validate:"required"`
	ThumbnailURL    string   `firestore:"thumbnailUrl"`
	BackgroundColor string   `firestore:"backgroundColor"`
	Labels          []string `firestore:"labels"`
}
