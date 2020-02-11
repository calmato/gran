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
		apiV1.POST("/boards", reg.V1Board.Create)

		apiV1.GET(":groupID/boards", reg.V1Board.Index)
	}

	return r
}
