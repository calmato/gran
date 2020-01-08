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

// CustomValidator - カスタムバリデーションインターフェース
type CustomValidator interface {
	Run(i interface{}) error
}

// customValidator - バリデーション用の構造体
type customValidator struct {
	validate *validator.Validate
}

// NewCustomValidator - Validatorの生成
func NewCustomValidator() CustomValidator {
	validate := validator.New()

	if err := validate.RegisterValidation("password", passwordCheck); err != nil {
		return nil
	}

	return &customValidator{
		validate: validate,
	}
}

// Run - バリデーションの実行
func (cv *customValidator) Run(i interface{}) error {
	return cv.validate.Struct(i)
}

func passwordCheck(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
}
