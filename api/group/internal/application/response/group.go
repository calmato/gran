package response

import "time"

// Group - グループのレスポンス
type Group struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	UserIDs       []string  `json:"user_ids,omitempty"`
	InvitedEmails []string  `json:"invitedEmails,omitempty"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `firestore:"created_at"`
	UpdatedAt     time.Time `firestore:"updated_at"`
}

// Groups - グループ一覧のレスポンス
type Groups struct {
	Groups []*Group `json:"groups"`
}
