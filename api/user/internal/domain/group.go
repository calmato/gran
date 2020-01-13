package domain

import "time"

// Group - Groupエンティティ
type Group struct {
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}
