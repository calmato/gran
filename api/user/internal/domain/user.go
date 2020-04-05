package domain

import "time"

// User - Userエンティティ
type User struct {
	ID           string    `firestore:"id"`
	Name         string    `firestore:"name"`
	DisplayName  string    `firestore:"display_name"`
	Email        string    `firestore:"email"`
	PhoneNumber  string    `firestore:"phone_number"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	Biography    string    `firestore:"biography"`
	Password     string    `firestore:"-"`
	GroupIDs     []string  `firestore:"group_ids,omitempty"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
