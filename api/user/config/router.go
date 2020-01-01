package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/registry"
)

// Router - ルーティングの定義
func Router() *gin.Engine {
	registry := registry.NewRegistry()

	// ルーティング
	r := gin.Default()

	// api v1 routes
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/", registry.APIV1HealthHandler.HealthCheck)
	}

	return r
}
