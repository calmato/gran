package validation

// バリデーションタグ
const (
	RequiredTag = "required"
	EqFieldTag  = "eqfield"
	MinTag      = "min"
	MaxTag      = "max"
	EmailTag    = "email"
	PasswordTag = "password"
)

// バリデーションメッセージ
const (
	RequiredDescription = "入力必須です"
	EqFieldDescription  = "%sと入力が一致しません"
	MinDescription      = "%s文字以上で入力してください"
	MaxDescription      = "%s文字以下で入力してください"
	EmailDescription    = "メールアドレスの形式で入力してください"
	PasswordDescription = "パスワードの形式で入力してください"
)

// カスタムバリデーションメッセージ
const (
	CustomUniqueDescription = "すでに存在します"
)
