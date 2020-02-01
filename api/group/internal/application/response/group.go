package response

import "time"

// Group - グループのレスポンス
type Group struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	UserRefs    []string  `json:"user_refs"`
	Description string    `json:"description"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

// Groups - グループ一覧のレスポンス
type Groups struct {
	Groups []*Group `json:"groups"`
}
