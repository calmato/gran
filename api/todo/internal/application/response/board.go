package response

import "time"

// Board - ボードのレスポンス
type Board struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	IsClosed        bool      `json:"isClosed"`
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

// TaskInShowBoard - ボード詳細用 タスクのレスポンス
type TaskInShowBoard struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Labels          []string  `json:"labels"`
	AssignedUserIDs []string  `json:"assignedUserIds"`
	DeadlinedAt     time.Time `json:"deadlinedAt"`
}

// ListInShowBoard - ボード詳細用 リストのレスポンス
type ListInShowBoard struct {
	ID    string             `json:"id"`
	Name  string             `json:"name"`
	Color string             `json:"color"`
	Tasks []*TaskInShowBoard `json:"tasks"`
}

// ShowBoard - ボード詳細のレスポンス
type ShowBoard struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	IsClosed        bool               `json:"isClosed"`
	ThumbnailURL    string             `json:"thumbnailUrl"`
	BackgroundColor string             `json:"backgroundColor"`
	Labels          []string           `json:"labels"`
	GroupID         string             `json:"groupId"`
	Lists           []*ListInShowBoard `json:"lists"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
}

// CreateBoardList - ボードリスト登録のレスポンス
type CreateBoardList struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	IsClosed        bool      `json:"isClosed"`
	ThumbnailURL    string    `json:"thumbnailUrl"`
	BackgroundColor string    `json:"backgroundColor"`
	Labels          []string  `json:"labels"`
	GroupID         string    `json:"groupId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
