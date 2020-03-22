package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/todo/internal/application"
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/interface/handler"
	"github.com/16francs/gran/api/todo/middleware"
)

// APIV1TaskHandler - Taskハンドラのインターフェース
type APIV1TaskHandler interface {
	Create(ctx *gin.Context)
}

type apiV1TaskHandler struct {
	taskApplication application.TaskApplication
}

// NewAPIV1TaskHandler - APIV1TaskHandlerの生成
func NewAPIV1TaskHandler(ta application.TaskApplication) APIV1TaskHandler {
	return &apiV1TaskHandler{
		taskApplication: ta,
	}
}

func (th *apiV1TaskHandler) Create(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")
	boardID := ctx.Params.ByName("boardID")

	req := &request.CreateTask{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := th.taskApplication.Create(c, groupID, boardID, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
