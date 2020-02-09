package validation

// バリデーションタグ
const (
	RequiredTag = "required"
	EqFieldTag  = "eqfield"
	MinTag      = "min"
	MaxTag      = "max"
	UniqueTag   = "unique"
)

// バリデーションメッセージ
const (
	RequiredMessage = "入力必須です"
	EqFieldMessage  = "%sと入力が一致しません"
	MinMessage      = "%s文字以上で入力してください"
	MaxMessage      = "%s文字以下で入力してください"
	UniqueMessage   = "重複する要素が存在します"
)

// カスタムバリデーションメッセージ
const (
	CustomUniqueMessage = "すでに存在します"
)
