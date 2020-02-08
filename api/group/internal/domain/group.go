package domain

import "time"

// Group - Groupエンティティ
type Group struct {
	ID            string    `firestore:"id"`
	Name          string    `firestore:"name"`
	Description   string    `firestore:"description"`
	UserRefs      []string  `firestore:"user_refs"`
	InvitedEmails []string  `firestore:"invited_emails"`
	CreatedAt     time.Time `firestore:"created_at"`
	UpdatedAt     time.Time `firestore:"updated_at"`
}
