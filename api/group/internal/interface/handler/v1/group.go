package v1

import (
	"net/http"

	"github.com/16francs/gran/api/group/internal/application/response"

	"github.com/16francs/gran/api/group/internal/application"
	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/interface/handler"
	"github.com/16francs/gran/api/group/middleware"
	"github.com/gin-gonic/gin"
)

// APIV1GroupHandler - Groupハンドラのインターフェース
type APIV1GroupHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
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

func (gh *apiV1GroupHandler) Index(ctx *gin.Context) {
	c := middleware.GinContextToContext(ctx)

	gs, err := gh.groupApplication.Index(c)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	gsr := make([]*response.Group, len(gs))
	for i, v := range gs {
		gr := &response.Group{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			UserRefs:    v.UserRefs,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}

		gsr[i] = gr
	}

	res := &response.Groups{
		Groups: gsr,
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
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		UserRefs:    g.UserRefs,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
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
