package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/todo/middleware"
	"github.com/16francs/gran/api/todo/registry"
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
	apiV1 := r.Group("/v1/todos")
	{
		groups := apiV1.Group("/groups/:groupID")
		{
			groups.GET("/boards", reg.V1Board.Index)
			groups.GET("/boards/:boardID", reg.V1Board.Show)
			groups.POST("/boards", reg.V1Board.Create)

			boards := groups.Group("/boards/:boardID")
			{
				boards.POST("/lists", reg.V1Board.CreateBoardList)
				boards.PATCH("/lists/:boardListID", reg.V1Board.UpdateBoardList)

				boards.POST("/tasks", reg.V1Task.Create)
				boards.PATCH("/kanban", reg.V1Board.UpdateKanban)
			}
		}

		apiV1.GET("/tasks/:taskID", reg.V1Task.Show)
	}

	return r
}
