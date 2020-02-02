package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/todo/registry"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry) *gin.Engine {
	// ルーティング
	r := gin.Default()

	// Corsの設定
	r.Use(SetCors())

	r.GET("/health", reg.Health.HealthCheck)

	// api v1 routes
	r.Group("/v1/todo")
	// apiV1 := r.Group("/v1/todo")

	return r
}
