package registry

import (
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	firebase.Firebase
	firestore.Firestore
	v1.APIV1HealthHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fb *firebase.Firebase, fs *firestore.Firestore) *Registry {
	apiV1HealthHandler := v1.NewAPIV1HealthHandler()

	return &Registry{
		Firebase:           *fb,
		Firestore:          *fs,
		APIV1HealthHandler: *apiV1HealthHandler,
	}
}
