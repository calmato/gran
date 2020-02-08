package domain

import "time"

// User - Userエンティティ
type User struct {
	ID           string    `firestore:"id"`
	Email        string    `firestore:"email" validate:"email"`
	Password     string    `firestore:"-" validate:"password"`
	Name         string    `firestore:"name"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	GroupRefs    []string  `firestore:"group_refs"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
