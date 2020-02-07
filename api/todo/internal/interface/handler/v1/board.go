package v1

import (
	"net/http"

	"github.com/16francs/gran/api/todo/internal/application"
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/interface/handler"
	"github.com/16francs/gran/api/todo/middleware"
	"github.com/gin-gonic/gin"
)

// APIV1BoardHandler - Boardハンドラのインターフェース
type APIV1BoardHandler interface {
	Create(ctx *gin.Context)
}

type apiV1BoardHandler struct {
	boardApplication application.BoardApplication
}

// NewAPIV1BoardHandler - APIV1BoardHandlerの生成
func NewAPIV1BoardHandler(ba application.BoardApplication) APIV1BoardHandler {
	return &apiV1BoardHandler{
		boardApplication: ba,
	}
}

func (bh *apiV1BoardHandler) Create(ctx *gin.Context) {
	req := request.CreateBoard{}
	if err := ctx.BindJSON(&req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := bh.boardApplication.Create(c, &req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
