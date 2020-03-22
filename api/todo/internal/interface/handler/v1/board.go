package v1

import (
	"net/http"

	"github.com/16francs/gran/api/todo/internal/application"
	"github.com/16francs/gran/api/todo/internal/application/request"
	"github.com/16francs/gran/api/todo/internal/application/response"
	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/interface/handler"
	"github.com/16francs/gran/api/todo/middleware"
	"github.com/gin-gonic/gin"
)

// APIV1BoardHandler - Boardハンドラのインターフェース
type APIV1BoardHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	CreateBoardList(ctx *gin.Context)
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

func (bh *apiV1BoardHandler) Index(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	c := middleware.GinContextToContext(ctx)

	bs, err := bh.boardApplication.Index(c, groupID)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	brs := make([]*response.Board, len(bs))
	for i, v := range bs {
		br := &response.Board{
			ID:              v.ID,
			Name:            v.Name,
			IsClosed:        v.IsClosed,
			ThumbnailURL:    v.ThumbnailURL,
			BackgroundColor: v.BackgroundColor,
			Labels:          append([]string{}, v.Labels...),
			GroupID:         v.GroupID,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		}

		brs[i] = br
	}

	res := &response.Boards{
		Boards: brs,
	}

	ctx.JSON(http.StatusOK, res)
}

func (bh *apiV1BoardHandler) Show(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")
	boardID := ctx.Params.ByName("boardID")

	c := middleware.GinContextToContext(ctx)

	b, err := bh.boardApplication.Show(c, groupID, boardID)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	res := &response.ShowBoard{
		ID:              b.ID,
		Name:            b.Name,
		IsClosed:        b.IsClosed,
		ThumbnailURL:    b.ThumbnailURL,
		BackgroundColor: b.BackgroundColor,
		Labels:          append([]string{}, b.Labels...),
		GroupID:         b.GroupID,
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
	}

	lrs := make([]*response.ListInShowBoard, len(b.Lists))
	for i, l := range b.Lists {
		lr := &response.ListInShowBoard{
			ID:    l.ID,
			Name:  l.Name,
			Color: l.Color,
		}

		trs := make([]*response.TaskInShowBoard, len(l.Tasks))
		for j, t := range l.Tasks {
			tr := &response.TaskInShowBoard{
				ID:              t.ID,
				Name:            t.Name,
				Labels:          append([]string{}, t.Labels...),
				AssignedUserIDs: append([]string{}, t.AssignedUserIDs...),
				DeadlinedAt:     t.DeadlinedAt,
			}

			trs[j] = tr
		}

		lr.Tasks = trs
		lrs[i] = lr
	}

	res.Lists = lrs

	ctx.JSON(http.StatusOK, res)
}

func (bh *apiV1BoardHandler) Create(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")

	req := &request.CreateBoard{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := bh.boardApplication.Create(c, groupID, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (bh *apiV1BoardHandler) CreateBoardList(ctx *gin.Context) {
	groupID := ctx.Params.ByName("groupID")
	boardID := ctx.Params.ByName("boardID")

	req := &request.CreateBoardList{}
	if err := ctx.BindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	c := middleware.GinContextToContext(ctx)
	if err := bh.boardApplication.CreateBoardList(c, groupID, boardID, req); err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
