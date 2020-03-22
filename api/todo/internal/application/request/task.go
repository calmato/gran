package request

import "time"

// CreateTask - タスク作成APIのリクエスト
type CreateTask struct {
	Name            string    `json:"name" label:"タスク名" validate:"required,max=32"`
	Description     string    `json:"description" label:"説明" validate:"max=256"`
	BoardListID     string    `json:"listId" label:"ボードリストID" validate:"required"`
	Labels          []string  `json:"labels" label:"ラベル一覧" validate:"unique,dive,max=16"`
	Attachments     []string  `json:"attachmentUrls" label:"添付ファイル一覧"`
	AssignedUserIDs []string  `json:"assignedUserIds" labels:"参加ユーザID一覧" validate:"unique"`
	DeadlinedAt     time.Time `json:"deadlinedAt" label:"期限"`
}
