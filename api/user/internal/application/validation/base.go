package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	passwordString = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$"
)

var (
	passwordRegex = regexp.MustCompile(passwordString)
)

// customValidator - バリデーション用の構造体
type customValidator struct {
	validate *validator.Validate
}

// NewCustomValidator - Validatorの生成
func NewCustomValidator() *customValidator {
	validate := validator.New()

	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return passwordRegex.MatchString(fl.Field().String())
	})

	return &customValidator{
		validate: validate,
	}
}

// Run - バリデーションの実行
func (v *customValidator) Run(i interface{}) error {
	return v.validate.Struct(i)
}
