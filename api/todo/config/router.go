package config

import (
	"github.com/gin-gonic/gin"

	"github.com/calmato/gran/api/todo/middleware"
	"github.com/calmato/gran/api/todo/registry"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry) *gin.Engine {
	// ルーティング
	r := gin.Default()

	// Corsの設定
	r.Use(SetCors())

	// Loggingの設定
	r.Use(middleware.Logging())

	r.GET("/health", reg.Health.HealthCheck)

	// api v1 routes
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/groups", reg.V1Group.Index)
		apiV1.GET("/groups/:groupID", reg.V1Group.Show)
		apiV1.POST("/groups", reg.V1Group.Create)
		apiV1.POST("/groups/:groupID/invite", reg.V1Group.InviteUsers)
		apiV1.POST("/groups/:groupID/join", reg.V1Group.Join)
		apiV1.PATCH("/groups/:groupID", reg.V1Group.Update)

		apiV1.GET("/tasks/:taskID", reg.V1Task.Show)

		groups := apiV1.Group("/groups/:groupID")
		{
			groups.GET("/boards", reg.V1Board.Index)
			groups.GET("/boards/:boardID", reg.V1Board.Show)
			groups.POST("/boards", reg.V1Board.Create)
			groups.PATCH("/boards/:boardID/kanban", reg.V1Board.UpdateKanban)

			groups.POST("/tasks", reg.V1Task.Create)

			boards := groups.Group("/boards/:boardID")
			{
				boards.POST("/lists", reg.V1Board.CreateBoardList)
				boards.PATCH("/lists/:boardListID", reg.V1Board.UpdateBoardList)
			}
		}
	}

	return r
}
