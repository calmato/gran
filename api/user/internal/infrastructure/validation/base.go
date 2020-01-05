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

// Validator - バリデーション用の構造体
type Validator struct {
	validate *validator.Validate
}

// NewValidator - Validatorの生成
func NewValidator() *Validator {
	validate := validator.New()

	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return passwordRegex.MatchString(fl.Field().String())
	})

	return &Validator{
		validate: validate,
	}
}
