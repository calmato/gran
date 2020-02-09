package request

// CreateUser - ユーザー登録APIのリクエスト
type CreateUser struct {
	Email                string `json:"email" label:"メールアドレス" validate:"required,email,max=256"`
	Password             string `json:"password" label:"パスワード" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" label:"パスワード(確認用)" validate:"required,eqfield=Password"`
}
