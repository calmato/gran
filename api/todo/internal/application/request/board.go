package request

import "time"

// CreateBoard - ボード作成APIのリクエスト
type CreateBoard struct {
	Name            string   `json:"name" label:"ボード名" validate:"required,max=64"`
	IsClosed        bool     `json:"isClosed" label:"非公開"`
	BackgroundColor string   `json:"backgroundColor" label:"背景色"`
	Thumbnail       string   `json:"thumbnail" label:"サムネイル"`
	Labels          []string `json:"labels" label:"ラベル一覧" validate:"unique,dive,max=16"`
}

// CreateBoardList - ボードリスト作成APIのリクエスト
type CreateBoardList struct {
	Name  string `json:"name" label:"ボードリスト名" validate:"required,max=64"`
	Color string `json:"color" label:"ボードリストの色"`
}

// UpdateBoardList - ボードリスト編集APIのリクエスト
type UpdateBoardList struct {
	Name  string `json:"name" label:"ボードリスト名" validate:"required,max=64"`
	Color string `json:"color" label:"ボードリストの色"`
}

// TaskInUpdateKanban - カンバン編集用 タスクのリクエスト
type TaskInUpdateKanban struct {
	ID              string    `json:"id" validate:"required"`
	Name            string    `json:"name"`
	Labels          []string  `json:"labels"`
	AssignedUserIDs []string  `json:"assignedUserIds"`
	DeadlinedAt     time.Time `json:"deadlinedAt"`
}

// ListInUpdateKanban - カンバン編集用 タスクのリクエスト
type ListInUpdateKanban struct {
	ID    string                `json:"id" validate:"required"`
	Name  string                `json:"name"`
	Color string                `json:"color"`
	Tasks []*TaskInUpdateKanban `json:"tasks"`
}

// UpdateKanban - カンバン編集APIのリクエスト
type UpdateKanban struct {
	Lists []*ListInUpdateKanban `json:"lists"`
}
