package domain

// ValidationError - バリデーションエラー用構造体
type ValidationError struct {
	Field       string
	Description string
}

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	ErrorCode ErrorCode
	Value     error
	Detail    interface{}
}

// ErrorCode - エラーの種類
type ErrorCode uint

const (
	// Unknown - 不明なエラー
	Unknown ErrorCode = iota
	// Unauthorized - 認証エラー
	Unauthorized
	// Forbidden - 権限エラー
	Forbidden
	// InvalidDomainValidation - ドメインのバリデーションエラー
	InvalidDomainValidation
	// InvalidRequestValidation - リクエストのバリデーションエラー
	InvalidRequestValidation
	// UnableParseJSON - JSON型から構造体への変換エラー
	UnableParseJSON
	// ErrorInDatastore - データストアでのエラー
	ErrorInDatastore
)

// ErrorCodeGetter - ErrorCodeを返すインターフェース
type ErrorCodeGetter interface {
	Type() ErrorCode
}

// ErrorDetailGetter - Detailを返すインターフェース
type ErrorDetailGetter interface {
	Show() interface{}
}

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error, detail ...interface{}) error {
	return CustomError{
		ErrorCode: ec,
		Value:     err,
		Detail:    detail,
	}
}

// Error - Errorを返すインターフェース
func (e CustomError) Error() string {
	return e.Value.Error()
}

// Type - ErrorCodeを返すインターフェース
func (e CustomError) Type() ErrorCode {
	return e.ErrorCode
}

// Show - エラー詳細を返すインターフェース
func (e CustomError) Show() interface{} {
	return e.Detail
}
