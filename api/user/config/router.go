package config

import (
	"github.com/gin-gonic/gin"

	"github.com/16francs/gran/api/user/registry"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry) *gin.Engine {
	// ルーティング
	r := gin.Default()

	// api v1 routes
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/", reg.APIV1HealthHandler.HealthCheck)

		apiV1.POST("/users", reg.APIV1UserHandler.Create)
	}

	return r
}
