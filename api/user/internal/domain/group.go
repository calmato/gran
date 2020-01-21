package domain

import "time"

// Group - Groupエンティティ
type Group struct {
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Users       []User    `firestore:"-"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}
