package response

import "time"

// Board - ボードのレスポンス
type Board struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Closed          bool      `json:"closed"`
	ThumbnailURL    string    `json:"thumbnailUrl"`
	BackgroundColor string    `json:"backgroundColor"`
	Labels          []string  `json:"labels"`
	GroupID         string    `json:"groupId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Boards - ボード一覧のレスポンス
type Boards struct {
	Boards []*Board `json:"boards"`
}
