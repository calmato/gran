package validation

import (
	"github.com/go-playground/validator/v10"
)

// FormValidator - リクエストフォームバリデーションインターフェース
type FormValidator interface {
	Run(i interface{}) error
}

// formValidator - バリデーション用の構造体
type formValidator struct {
	validate *validator.Validate
}

// NewFormValidator - Validatorの生成
func NewFormValidator() FormValidator {
	validate := validator.New()

	return &formValidator{
		validate: validate,
	}
}

// Run - バリデーションの実行
func (fv *formValidator) Run(i interface{}) error {
	return fv.validate.Struct(i)
}
