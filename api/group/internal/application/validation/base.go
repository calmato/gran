package validation

import (
	"github.com/go-playground/validator/v10"
)

// RequestValidator - リクエストバリデーションインターフェース
type RequestValidator interface {
	Run(i interface{}) error
}

type requestValidator struct {
	validate validator.Validate
}

// NewRequestValidator - Validatorの生成
func NewRequestValidator() RequestValidator {
	validate := validator.New()

	return &requestValidator{
		validate: *validate,
	}
}

// Run - バリデーションの実行
func (rv *requestValidator) Run(i interface{}) error {
	return rv.validate.Struct(i)
}
