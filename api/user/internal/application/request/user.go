package request

// CreateUser - ユーザー登録APIのリクエスト
type CreateUser struct {
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// CreateGroup - グループ登録APIのリクエスト
type CreateGroup struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
