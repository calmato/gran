package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/group/internal/application/response"
	"github.com/16francs/gran/api/group/internal/domain"
)

// ErrorHandling - エラーレスポンスを返す
func ErrorHandling(ctx *gin.Context, err error) {
	res := errorResponse(err)

	ctx.JSON(res.StatusCode, res)
	ctx.Abort()
}

// errorResponse - エラー用のレスポンスを返す
func errorResponse(err error) *response.ErrorResponse {
	var res *response.ErrorResponse

	switch errorCode(err) {
	case domain.InvalidDomainValidation:
		res = response.BadRequest
		log.Printf("info: BadRequest: %v", err.Error())
		res.Description = "" // TODO: バリデーションエラーの結果入れる
	case domain.InvalidRequestValidation:
		res = response.BadRequest
		log.Printf("info: BadRequest: %v", err.Error())
		res.Description = "" // TODO: バリデーションエラーの結果入れる
	case domain.Unauthorized:
		log.Printf("info: Unauthorized: %v", err.Error())
		res = response.Unauthorized
	case domain.Forbidden:
		log.Printf("info: Forbidden: %v", err.Error())
		res = response.Forbidden
	default:
		log.Printf("error: Internal Server Error: %v", err.Error())
		res = response.InternalServerError
	}

	res.ErrorCode = errorCode(err)
	return res
}

// errorCode - ErrorCodeを持つ場合はそれを返し、無ければUnknownを返す
func errorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.ErrorCodeGetter); ok {
		return e.Type()
	}

	return domain.Unknown
}
