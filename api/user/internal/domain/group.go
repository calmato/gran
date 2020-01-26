package domain

import "time"

// Group - Groupエンティティ
type Group struct {
	ID          string    `firestore:"-"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	UserRefs    []string  `firestore:"user_refs"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}
