package response

import "time"

// Group - グループのレスポンス
type Group struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	UserIDs       []string  `json:"userIds,omitempty"`
	BoardIDs      []string  `json:"boardIds,omitempty"`
	InvitedEmails []string  `json:"invitedEmails,omitempty"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
}

// Groups - グループ一覧のレスポンス
type Groups struct {
	Groups []*Group `json:"groups"`
}
