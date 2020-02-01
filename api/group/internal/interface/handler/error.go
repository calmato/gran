package handler

import (
	"encoding/json"
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

	switch getErrorCode(err) {
	case domain.InvalidDomainValidation:
		res = response.BadRequest
		setValidationErrors(res, err)
		logging("info", "BadRequest", err, res.ValidationErrors...)
	case domain.InvalidRequestValidation:
		res = response.BadRequest
		setValidationErrors(res, err)
		logging("info", "BadRequest", err, res.ValidationErrors...)
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

	res.ErrorCode = getErrorCode(err)
	return res
}

func logging(level string, message string, err error, ves ...*response.ValidationError) {
	log.Printf("%s: %s: %v", level, message, err.Error())

	if len(ves) > 0 {
		j, _ := json.Marshal(ves)

		log.Printf("%s: %s", level, j)
	}
}

func getErrorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.ShowError); ok {
		return e.Code()
	}

	return domain.Unknown
}

func setValidationErrors(er *response.ErrorResponse, err error) {
	if e, ok := err.(domain.ShowError); ok {
		ves := e.Validation()
		er.ValidationErrors = make([]*response.ValidationError, len(ves))

		for i, ve := range ves {
			er.ValidationErrors[i] = &response.ValidationError{
				Field:   ve.Field,
				Message: ve.Message,
			}
		}
	}
}
