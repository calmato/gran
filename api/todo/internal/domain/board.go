package domain

import "time"

// Board - Boardエンティティ
type Board struct {
	ID              string    `firestore:"id"`
	Name            string    `firestore:"name"`
	Closed          bool      `firestore:"closed"`
	ThumbnailURL    string    `firestore:"thumbnail_url"`
	BackgroundColor string    `firestore:"background_color"`
	Labels          []string  `firestore:"labels"`
	GroupID         string    `firestore:"group_id"`
	ListIDs         []string  `firestore:"list_ids"`
	CreatedAt       time.Time `firestore:"created_at"`
	UpdatedAt       time.Time `firestore:"updated_at"`
}

// BoardList - BoardListエンティティ
type BoardList struct {
	ID        string    `firestore:"id"`
	Name      string    `firestore:"name"`
	Color     string    `firestore:"color"`
	BoardID   string    `firestore:"board_id"`
	TaskIDs   []string  `firestore:"task_ids"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
