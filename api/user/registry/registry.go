package registry

import (
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	authentication.Auth
	firestore.Firestore
	v1.APIV1HealthHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	apiV1HealthHandler := v1HealthInjection()

	return &Registry{
		Auth:               *fa,
		Firestore:          *fs,
		APIV1HealthHandler: apiV1HealthHandler,
	}
}

func v1HealthInjection() v1.APIV1HealthHandler {
	apiV1HealthHandler := v1.NewAPIV1HealthHandler()

	return apiV1HealthHandler
}
