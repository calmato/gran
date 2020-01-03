package domain

import "time"

// User - Userエンティティ
type User struct {
	ID           string    `firestore:"id"`
	Email        string    `firestore:"email"`
	Password     string    `firestore:"-"`
	Name         string    `firestore:"name"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
