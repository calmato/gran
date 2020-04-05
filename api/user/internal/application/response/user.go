package response

import "time"

// ShowProfile - ログインユーザー取得APIのレスポンス
type ShowProfile struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	DisplayName  string    `json:"displayName"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Biography    string    `json:"biography"`
	GroupIDs     []string  `json:"groupIds,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
