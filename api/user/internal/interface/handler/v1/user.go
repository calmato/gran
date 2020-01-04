package v1

import (
	"net/http"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type APIV1UserHandler interface {
	Create(ctx *gin.Context)
}

type apiV1UserHandler struct {
	usecase usecase.UserUsecase
}

func NewAPIV1UserHandler(uu usecase.UserUsecase) APIV1UserHandler {
	return &apiV1UserHandler{
		usecase: uu,
	}
}

func (uh *apiV1UserHandler) Create(ctx *gin.Context) {
	req := request.CreateUser{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{}) // TODO: エラーハンドラ作成
		return
	}

	if err := uh.usecase.Create(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{}) // TODO: エラーハンドラ作成
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
