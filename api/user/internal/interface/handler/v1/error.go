package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/internal/domain"
)

func ErrorHandling(ctx *gin.Çontext, err error) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"status":      http.StatusText(http.StatusForbidden),
		"description": "認証エラーです．",
	})

	ctx.Abort()
}

func statusCode(err error) uint {
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
