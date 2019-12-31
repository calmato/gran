package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIV1HealthHandler - ヘルスチェック用のハンドラ
type APIV1HealthHandler struct{}

// NewAPIV1HealthHandler - APIV1HealthHandlerの生成
func NewAPIV1HealthHandler() *APIV1HealthHandler {
	return &APIV1HealthHandler{}
}

// HealthCheck - ヘルスチェック
func (h *APIV1HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
