package request

// CreateUser - ユーザー登録APIのリクエスト
type CreateUser struct {
	Email                string `json:"email" validate:"required,email,max=256"`
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}
