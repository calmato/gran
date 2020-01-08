package response

import (
	"net/http"

	"github.com/16francs/gran/api/user/internal/domain"
)

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode  int              `json:"statusCode"`
	ErrorCode   domain.ErrorCode `json:"errorCode"`
	Message     string           `json:"message"`
	Description interface{}      `json:"description"`
}

var BadRequest = &ErrorResponse{
	StatusCode: http.StatusBadRequest,
	Message:    "不正なパラメータが入力されています。",
}

var Unauthorized = &ErrorResponse{
	StatusCode:  http.StatusUnauthorized,
	Message:     "認証に必要な情報がありません。",
	Description: "",
}

var Forbidden = &ErrorResponse{
	StatusCode:  http.StatusForbidden,
	Message:     "その操作を実行する権限がありません。",
	Description: "",
}

var InternalServerError = &ErrorResponse{
	StatusCode:  http.StatusInternalServerError,
	Message:     "異常な処理が検出されました。",
	Description: "",
}
