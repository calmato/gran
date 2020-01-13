package v1

import (
	"net/http"

	"github.com/16francs/gran/api/user/internal/application"
	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/interface/handler"
	"github.com/gin-gonic/gin"
)

// APIV1UserHandler - Userハンドラのインターフェース
type APIV1UserHandler interface {
	Create(ctx *gin.Context)
	CreateGroup(ctx *gin.Context)
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
	req := request.CreateUser{}
	if err := ctx.BindJSON(&req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	if err := uh.userApplication.Create(ctx, &req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (uh *apiV1UserHandler) CreateGroup(ctx *gin.Context) {
	// TODO: 認証情報の取得

	req := request.CreateGroup{}
	if err := ctx.BindJSON(&req); err != nil {
		// TODO: エラー処理
		handler.ErrorHandling(ctx, err)
		return
	}

	if err := uh.userApplication.CreateGroup(ctx, &req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
