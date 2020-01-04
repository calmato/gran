package input

// CreateUser - ユーザー登録APIのリクエスト
type CreateUser struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
