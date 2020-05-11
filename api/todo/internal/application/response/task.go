package response

import "time"

// ShowTask - タスク詳細のレスポンス
type ShowTask struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Labels          []string  `json:"labels"`
	AttachmentURLs  []string  `json:"attachmentUrls"`
	GroupID         string    `json:"groupID"`
	BoardID         string    `json:"boardId"`
	AssignedUserIDs []string  `json:"assignedUserIds"`
	CheckListIDs    []string  `json:"checklistIds"`
	CommentIDs      []string  `json:"commentIds"`
	DeadlinedAt     time.Time `json:"deadlinedAt"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// CreateTask - タスク詳細のレスポンス
type CreateTask struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Labels          []string  `json:"labels"`
	AttachmentURLs  []string  `json:"attachmentUrls"`
	GroupID         string    `json:"groupID"`
	BoardID         string    `json:"boardId"`
	AssignedUserIDs []string  `json:"assignedUserIds"`
	DeadlinedAt     time.Time `json:"deadlinedAt"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
