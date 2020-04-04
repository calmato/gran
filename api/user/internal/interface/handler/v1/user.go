package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/internal/application"
	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/interface/handler"
	"github.com/16francs/gran/api/user/middleware"
)

// APIV1UserHandler - Userハンドラのインターフェース
type APIV1UserHandler interface {
	Create(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type apiV1UserHandler struct {
	userApplication application.UserApplication
}

// NewAPIV1UserHandler - APIV1UserHandlerの生成
func NewAPIV1UserHandler(ua application.UserApplication) APIV1UserHandler {
	return &apiV1UserHandler{
		userApplication: ua,
	}
}

func (uh *apiV1UserHandler) Create(ctx *gin.Context) {
	req := &request.CreateUser{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	if err := uh.userApplication.Create(ctx, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (uh *apiV1UserHandler) UpdateProfile(ctx *gin.Context) {
	c := middleware.GinContextToContext(ctx)

	req := &request.UpdateProfile{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	if err := uh.userApplication.UpdateProfile(c, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
