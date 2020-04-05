package domain

import "time"

// Task - Taskエンティティ
type Task struct {
	ID              string    `firestore:"id"`
	Name            string    `firestore:"name"`
	Description     string    `firestore:"description"`
	Labels          []string  `firestore:"labels"`
	AttachmentURLs  []string  `firestore:"attachment_urls"`
	GroupID         string    `firestore:"group_id"`
	BoardID         string    `firestore:"board_id"`
	BoardListID     string    `firestore:"-"`
	AssignedUserIDs []string  `firestore:"assigned_user_ids"`
	CheckListIDs    []string  `firestore:"checklist_ids"`
	CommentIDs      []string  `firestore:"comment_ids"`
	DeadlinedAt     time.Time `firestore:"deadlined_at"`
	CreatedAt       time.Time `firestore:"created_at"`
	UpdatedAt       time.Time `firestore:"updated_at"`
}
