package registry

import (
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
)

// Registry - DIコンテナ
type Registry struct {
	v1.APIV1HealthHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry() *Registry {
	apiV1HealthHandler := v1.NewAPIV1HealthHandler()

	return &Registry{
		*apiV1HealthHandler,
	}
}
