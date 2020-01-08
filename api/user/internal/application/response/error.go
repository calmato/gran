package response

import (
	"github.com/16francs/gran/api/user/internal/domain"
)

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode  int              `json:"statusCode"`
	ErrorCode   domain.ErrorCode `json:"errorCode"`
	Message     string           `json:"message"`
	Description interface{}      `json:"description"`
}
