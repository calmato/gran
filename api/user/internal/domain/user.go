package domain

import "time"

// User - Userエンティティ
type User struct {
	ID           string    `firestore:"id"`
	Email        string    `firestore:"email" validate:"email,max=256"`
	Password     string    `firestore:"-" validate:"password,min=6,max=32"`
	Name         string    `firestore:"name" validate:"max=32"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
