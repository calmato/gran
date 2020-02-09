package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/validation"
)

const (
	passwordString = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$"
)

var (
	passwordRegex = regexp.MustCompile(passwordString)
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

	if err := validate.RegisterValidation("password", passwordCheck); err != nil {
		return nil
	}

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

func passwordCheck(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
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
	case validation.EmailTag:
		return validation.EmailMessage
	case validation.PasswordTag:
		return validation.PasswordMessage
	default:
		return ""
	}
}
