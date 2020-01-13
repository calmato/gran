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

// DomainValidator - ドメインバリデーションインターフェース
type DomainValidator interface {
	Run(i interface{}) error
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
func (dv *domainValidator) Run(i interface{}) error {
	return dv.validate.Struct(i)
}

func passwordCheck(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
}
