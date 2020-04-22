package response

import "time"

// Group - グループのレスポンス
type Group struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	UserIDs       []string  `json:"userIds"`
	BoardIDs      []string  `json:"boardIds"`
	InvitedEmails []string  `json:"invitedEmails"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
}

// Groups - グループ一覧のレスポンス
type Groups struct {
	Groups []*Group `json:"groups"`
}
