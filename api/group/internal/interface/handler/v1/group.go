package v1

import (
	"net/http"

	"github.com/16francs/gran/api/group/internal/application"
	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/interface/handler"
	"github.com/16francs/gran/api/group/middleware"
	"github.com/gin-gonic/gin"
)

// APIV1GroupHandler - Groupハンドラのインターフェース
type APIV1GroupHandler interface {
	Create(ctx *gin.Context)
}

type apiV1GroupHandler struct {
	groupApplication application.GroupApplication
}

// NewAPIV1GroupHandler - APIV1GroupHandlerの生成
func NewAPIV1GroupHandler(ga application.GroupApplication) APIV1GroupHandler {
	return &apiV1GroupHandler{
		groupApplication: ga,
	}
}

func (gh *apiV1GroupHandler) Create(ctx *gin.Context) {
	req := request.CreateGroup{}
	if err := ctx.BindJSON(&req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := gh.groupApplication.Create(c, &req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
