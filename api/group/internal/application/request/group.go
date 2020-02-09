package request

// CreateGroup - グループ登録APIのリクエスト
type CreateGroup struct {
	Name        string `json:"name" label:"グループ名" validate:"required"`
	Description string `json:"description" label:"説明" validate:"max=256"`
}

// UpdateGroup - グループ編集APIのリクエスト
type UpdateGroup struct {
	Name        string `json:"name" label:"グループ名" validate:"required"`
	Description string `json:"description" label:"説明" validate:"max=256"`
}
