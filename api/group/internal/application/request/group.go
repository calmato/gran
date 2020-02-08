package request

// CreateGroup - グループ登録APIのリクエスト
type CreateGroup struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// UpdateGroup - グループ編集APIのリクエスト
type UpdateGroup struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
