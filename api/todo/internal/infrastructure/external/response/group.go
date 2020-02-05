package response

import "time"

// Group - グループのレスポンス
type Group struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	UserRefs    []string  `json:"user_refs"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
