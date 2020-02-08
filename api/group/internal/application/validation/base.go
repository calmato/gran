package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/validation"
)

// RequestValidator - リクエストバリデーションインターフェース
type RequestValidator interface {
	Run(i interface{}) []*domain.ValidationError
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
func (rv *requestValidator) Run(i interface{}) []*domain.ValidationError {
	err := rv.validate.Struct(i)
	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)
	validationErrors := make([]*domain.ValidationError, len(errors))

	for i, v := range errors {
		validationErrors[i] = &domain.ValidationError{
			Field:   v.Field(),
			Message: validationMessage(v.Tag(), v.Param()),
		}
	}

	return validationErrors
}

func validationMessage(tag string, param string) string {
	switch tag {
	case validation.RequiredTag:
		return validation.RequiredMessage
	case validation.EqFieldTag:
		return fmt.Sprintf(validation.EqFieldMessage, param)
	case validation.MinTag:
		return fmt.Sprintf(validation.MinMessage, param)
	case validation.MaxTag:
		return fmt.Sprintf(validation.MaxMessage, param)
	default:
		return ""
	}
}
