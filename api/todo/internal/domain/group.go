package domain

import "time"

// Group - Groupエンティティ
type Group struct {
	ID            string            `firestore:"id"`
	Name          string            `firestore:"name"`
	Description   string            `firestore:"description"`
	UserIDs       []string          `firestore:"user_ids"`
	BoardIDs      []string          `firestore:"-"`
	Boards        map[string]*Board `firestore:"-"`
	InvitedEmails []string          `firestore:"invited_emails,omitempty"`
	CreatedAt     time.Time         `firestore:"created_at"`
	UpdatedAt     time.Time         `firestore:"updated_at"`
}
