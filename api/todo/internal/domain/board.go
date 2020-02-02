package domain

import "time"

// Board - Boardエンティティ
type Board struct {
	ID              string    `firestore:"id"`
	Name            string    `firestore:"name" validate:"max=64"`
	Closed          bool      `firestore:"closed"`
	ThumbnailURL    string    `firestore:"thumbnail_url"`
	BackgroundColor string    `firestore:"background_url"`
	Labels          []string  `firestore:"labels"`
	BoardRef        string    `firestore:"board_ref"`
	CreatedAt       time.Time `firestore:"created_at"`
	UpdatedAt       time.Time `firestore:"updated_at"`
}
