package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/internal/application/response"
	"github.com/16francs/gran/api/user/internal/domain"
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
		res.Description = errorDetail(err)
		logging("info", "BadRequest", err)
	case domain.InvalidRequestValidation:
		res = response.BadRequest
		res.Description = errorDetail(err)
		logging("info", "BadRequest", err)
	case domain.Unauthorized:
		res = response.Unauthorized
		logging("info", "Unauthorized", err)
	case domain.Forbidden:
		res = response.Forbidden
		logging("info", "Forbidden", err)
	case domain.ErrorInDatastore:
		res = response.InternalServerError
		logging("error", "Internal Server Error", err)
	default:
		res = response.InternalServerError
		logging("error", "Internal Server Error", err)
	}

	res.ErrorCode = errorCode(err)
	return res
}

func logging(level string, message string, err error) {
	log.Printf("%s: %s: %v", level, message, err.Error())

	// バリデーションエラーの時、エラーレスポンスも出力
	if ves := errorDetail(err); len(ves) > 0 {
		for _, v := range ves {
			log.Printf("debug: - %s ->%s", v.Field, v.Description)
		}
	}
}

func errorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.ErrorCodeGetter); ok {
		return e.Type()
	}

	return domain.Unknown
}

func errorDetail(err error) []*domain.ValidationError {
	if e, ok := err.(domain.ValidationErrorGetter); ok {
		return e.Show()
	}

	return make([]*domain.ValidationError, 0)
}
