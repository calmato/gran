package validation

import (
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

// DomainValidator - ドメインバリデーションインターフェース
type DomainValidator interface {
	Run(i interface{}) []*domain.ValidationError
}

type domainValidator struct {
	validate validator.Validate
}

// NewDomainValidator - Validatorの生成
func NewDomainValidator() DomainValidator {
	validate := validator.New()

	if err := validate.RegisterValidation("password", passwordCheck); err != nil {
		return nil
	}

	return &domainValidator{
		validate: *validate,
	}
}

// Run - バリデーションの実行
func (dv *domainValidator) Run(i interface{}) []*domain.ValidationError {
	err := dv.validate.Struct(i)
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

func validationMessage(tag string) string {
	switch tag {
	case validation.EmailTag:
		return validation.EmailMessage
	case validation.PasswordTag:
		return validation.PasswordMessage
	default:
		return ""
	}
}
