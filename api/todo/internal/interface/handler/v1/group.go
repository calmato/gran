package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/calmato/gran/api/todo/internal/application"
	"github.com/calmato/gran/api/todo/internal/application/request"
	"github.com/calmato/gran/api/todo/internal/application/response"
	"github.com/calmato/gran/api/todo/internal/domain"
	"github.com/calmato/gran/api/todo/internal/interface/handler"
	"github.com/calmato/gran/api/todo/middleware"
)

// APIV1GroupHandler - Groupハンドラのインターフェース
type APIV1GroupHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	InviteUsers(ctx *gin.Context)
	Join(ctx *gin.Context)
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

func (gh *apiV1GroupHandler) Index(ctx *gin.Context) {
	c := middleware.GinContextToContext(ctx)

	gs, err := gh.groupApplication.Index(c)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	grs := make([]*response.Group, len(gs))
	for i, v := range gs {
		gr := &response.Group{
			ID:            v.ID,
			Name:          v.Name,
			Description:   v.Description,
			BoardIDs:      append([]string{}, v.BoardIDs...),
			UserIDs:       append([]string{}, v.UserIDs...),
			InvitedEmails: append([]string{}, v.InvitedEmails...),
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		}

		grs[i] = gr
	}

	res := &response.Groups{
		Groups: grs,
	}

	ctx.JSON(http.StatusOK, res)
}

func (gh *apiV1GroupHandler) Show(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	c := middleware.GinContextToContext(ctx)

	g, err := gh.groupApplication.Show(c, groupID)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	res := &response.Group{
		ID:            g.ID,
		Name:          g.Name,
		Description:   g.Description,
		UserIDs:       g.UserIDs,
		InvitedEmails: g.InvitedEmails,
		CreatedAt:     g.CreatedAt,
		UpdatedAt:     g.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

func (gh *apiV1GroupHandler) Create(ctx *gin.Context) {
	req := &request.CreateGroup{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := gh.groupApplication.Create(c, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (gh *apiV1GroupHandler) Update(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	req := &request.UpdateGroup{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)

	err := gh.groupApplication.Update(c, groupID, req)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}
}

func (gh *apiV1GroupHandler) InviteUsers(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	req := &request.InviteUsers{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)

	err := gh.groupApplication.InviteUsers(c, groupID, req)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}
}

func (gh *apiV1GroupHandler) Join(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	c := middleware.GinContextToContext(ctx)

	err := gh.groupApplication.Join(c, groupID)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}
}
