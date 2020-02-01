package domain

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	ErrorCode ErrorCode
	Value     error
}

// ErrorCodeGetter - ErrorCodeを返すインターフェース
type ErrorCodeGetter interface {
	Type() ErrorCode
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
)

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error) error {
	return CustomError{
		ErrorCode: ec,
		Value:     err,
	}
}

// Error - errorインターフェースを実装する
func (e CustomError) Error() string {
	return e.Value.Error()
}

// Type - ErrorCodeGetterインターフェースを実装する
func (e CustomError) Type() ErrorCode {
	return e.ErrorCode
}
