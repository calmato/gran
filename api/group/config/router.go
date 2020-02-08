package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/group/registry"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry) *gin.Engine {
	// ルーティング
	r := gin.Default()

	// Corsの設定
	r.Use(SetCors())

	r.GET("/health", reg.Health.HealthCheck)

	// api v1 routes
	apiV1 := r.Group("/v1/groups")
	{
		apiV1.GET("", reg.V1Group.Index)
		apiV1.POST("", reg.V1Group.Create)

		apiV1.GET("/:groupID", reg.V1Group.Show)
		apiV1.PATCH("/:groupID", reg.V1Group.Update)
	}

	return r
}
