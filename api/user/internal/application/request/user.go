package request

// CreateUser - ユーザー登録APIのリクエスト
type CreateUser struct {
	Email                string `json:"email" validate:"required,max=256"`
	Password             string `json:"password" validate:"required,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,max=32,eqfield=Password"`
}
