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

// BoardsInIndexGroup - グループ一覧用 ボード一覧レスポンス
type BoardsInIndexGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GroupInIndexGroup - グループ一覧用 グループのレスポンス
type GroupInIndexGroup struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	UserIDs       []string              `json:"userIds"`
	Boards        []*BoardsInIndexGroup `json:"boards"`
	InvitedEmails []string              `json:"invitedEmails"`
	Description   string                `json:"description"`
	CreatedAt     time.Time             `firestore:"createdAt"`
	UpdatedAt     time.Time             `firestore:"updatedAt"`
}

// IndexGroups - グループ一覧のレスポンス
type IndexGroups struct {
	Groups []*GroupInIndexGroup `json:"groups"`
}
