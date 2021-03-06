package domain

// ErrorCode - エラーの種類
type ErrorCode uint

const (
	// Unknown - 不明なエラー
	Unknown ErrorCode = iota
	// Unauthorized - 認証エラー
	Unauthorized
	// Forbidden - 権限エラー
	Forbidden
	// NotFound - 取得エラー
	NotFound
	// InvalidDomainValidation - ドメインのバリデーションエラー
	InvalidDomainValidation
	// InvalidRequestValidation - リクエストのバリデーションエラー
	InvalidRequestValidation
	// UnableParseJSON - JSON型から構造体への変換エラー
	UnableParseJSON
	// ErrorInDatastore - データストアでのエラー
	ErrorInDatastore
	// ErrorInStorage - ストレージでのエラー
	ErrorInStorage
	// AlreadyExists - ユニークチェックでのエラー
	AlreadyExists
	// NotEqualRequestWithDatastore - リクエスト値がデータストアの値と一致しない
	NotEqualRequestWithDatastore
)

// ShowError - エラー内容を返すインターフェース
type ShowError interface {
	Code() ErrorCode
	Error() string
	Validation() []*ValidationError
}

// ValidationError - バリデーションエラー用構造体
type ValidationError struct {
	Field   string
	Message string
}

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	ErrorCode        ErrorCode
	Value            error
	ValidationErrors []*ValidationError
}

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error, ves ...*ValidationError) error {
	return CustomError{
		ErrorCode:        ec,
		Value:            err,
		ValidationErrors: ves,
	}
}

// Code - エラーコードを返す
func (ce CustomError) Code() ErrorCode {
	return ce.ErrorCode
}

// Error - エラー内容を返す
func (ce CustomError) Error() string {
	return ce.Value.Error()
}

// Validation - エラー詳細を返す
func (ce CustomError) Validation() []*ValidationError {
	return ce.ValidationErrors
}
