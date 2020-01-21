package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/registry"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry) *gin.Engine {
	// ルーティング
	r := gin.Default()

	// Corsの設定
	r.Use(SetCors())

	r.GET("/health", reg.V1Health.HealthCheck)

	// api v1 routes
	apiV1 := r.Group("/v1/users")
	{
		apiV1.POST("", reg.V1User.Create)
		apiV1.POST("/group", reg.V1Group.Create)
	}

	return r
}
