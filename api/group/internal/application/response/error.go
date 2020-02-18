package response

import (
	"net/http"

	"github.com/16francs/gran/api/group/internal/domain"
)

// ValidationError - バリデーションエラーのレスポンス
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode       int                `json:"status"`
	ErrorCode        domain.ErrorCode   `json:"code"`
	Message          string             `json:"message"`
	ValidationErrors []*ValidationError `json:"errors,omitempty"`
}

// ステータスコードを付与したエラーレスポンス
var (
	BadRequest = &ErrorResponse{
		StatusCode: http.StatusBadRequest, // 404
		Message:    "不正なパラメータが入力されています。",
	}

	Unauthorized = &ErrorResponse{
		StatusCode:       http.StatusUnauthorized, // 401
		Message:          "認証に必要な情報がありません。",
		ValidationErrors: nil,
	}

	Forbidden = &ErrorResponse{
		StatusCode:       http.StatusForbidden, // 403
		Message:          "その操作を実行する権限がありません。",
		ValidationErrors: nil,
	}

	AlreadyExists = &ErrorResponse{
		StatusCode: http.StatusConflict, // 409
		Message:    "不正なパラメータが入力されています。",
	}

	InternalServerError = &ErrorResponse{
		StatusCode:       http.StatusInternalServerError, // 500
		Message:          "異常な処理が検出されました。",
		ValidationErrors: nil,
	}
)
