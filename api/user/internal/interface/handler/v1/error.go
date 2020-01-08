package v1

import (
	"net/http"

	"github.com/16francs/gran/api/user/internal/application/response"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/internal/domain"
)

// ErrorHandling - エラーレスポンスを返す
func ErrorHandling(ctx *gin.Context, err error) {
	res := &response.ErrorResponse{
		StatusCode:  statusCode(err),
		ErrorCode:   errorCode(err),
		Message:     "",
		Description: "",
	}

	ctx.JSON(res.StatusCode, res)
	ctx.Abort()
}

// statusCode - HTTPのステータスコードを取得
func statusCode(err error) int {
	switch errorCode(err) {
	case domain.InvalidDomainValidation:
		return http.StatusBadRequest // 400
	case domain.InvalidRequestValidation:
		return http.StatusBadRequest // 400
	case domain.Unauthorized:
		return http.StatusUnauthorized // 401
	case domain.Forbidden:
		return http.StatusForbidden // 403
	default:
		return http.StatusInternalServerError // 500
	}
}

// errorCode - ErrorCodeを持つ場合はそれを返し、無ければUnknownを返す
func errorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.ErrorCodeGetter); ok {
		return e.Type()
	}

	return domain.Unknown
}
