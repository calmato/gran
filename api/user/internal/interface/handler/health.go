package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIV1HealthHandler - ヘルスチェック
type APIV1HealthHandler interface {
	HealthCheck(ctx *gin.Context)
}

// apiV1HealthHandler - APIV1HealthCheckハンドラ
type apiV1HealthHandler struct {
}

// NewAPIV1HealthHandler - apiV1HealthHandlerの生成
func NewAPIV1HealthHandler() APIV1HealthHandler {
	return &apiV1HealthHandler{}
}

// HealthCheck - ヘルスチェック
func (hh *apiV1HealthHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
