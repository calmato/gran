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

// ShowBoard - ボード詳細のレスポンス
type ShowBoard struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Closed          bool     `json:"closed"`
	ThumbnailURL    string   `json:"thumbnailUrl"`
	BackgroundColor string   `json:"backgroundColor"`
	Labels          []string `json:"labels"`
	GroupID         string   `json:"groupId"`
	Lists           []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
		Tasks []struct {
			ID              string    `json:"id"`
			Name            string    `json:"name"`
			Labels          []string  `json:"labels"`
			AssignedUserIDs []string  `json:"assignedUserIds"`
			DeadlinedAt     time.Time `json:"deadlinedAt"`
		} `json:"tasks"`
	} `json:"lists"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
