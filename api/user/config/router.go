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

	// TODO: Logger(Middleware)の設定

	// api v1 routes
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/", registry.APIV1HealthHandler.HealthCheck)
	}

	// TODO: r.NoRouteの設定

	return r
}
