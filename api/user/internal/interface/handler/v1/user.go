package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/calmato/gran/api/user/internal/application"
	"github.com/calmato/gran/api/user/internal/application/request"
	"github.com/calmato/gran/api/user/internal/application/response"
	"github.com/calmato/gran/api/user/internal/domain"
	"github.com/calmato/gran/api/user/internal/interface/handler"
	"github.com/calmato/gran/api/user/middleware"
)

// APIV1UserHandler - Userハンドラのインターフェース
type APIV1UserHandler interface {
	Create(ctx *gin.Context)
	ShowProfile(ctx *gin.Context)
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

func (uh *apiV1UserHandler) ShowProfile(ctx *gin.Context) {
	c := middleware.GinContextToContext(ctx)

	u, err := uh.userApplication.ShowProfile(c)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	res := &response.ShowProfile{
		ID:           u.ID,
		Name:         u.Name,
		DisplayName:  u.DisplayName,
		Email:        u.Email,
		PhoneNumber:  u.PhoneNumber,
		ThumbnailURL: u.ThumbnailURL,
		Biography:    u.Biography,
		GroupIDs:     u.GroupIDs,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
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
