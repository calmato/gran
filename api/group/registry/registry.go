package registry

import (
	"github.com/16francs/gran/api/group/internal/interface/handler"
	v1 "github.com/16francs/gran/api/group/internal/interface/handler/v1"
	"github.com/16francs/gran/api/group/lib/firebase/authentication"
	"github.com/16francs/gran/api/group/lib/firebase/firestore"
	"github.com/16francs/gran/api/group/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	Health  handler.APIHealthHandler
	V1Group v1.APIV1GroupHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore, _ *storage.Storage) *Registry {
	health := healthInjection()
	v1Group := V1GroupInjection(fa, fs)

	return &Registry{
		Health:  health,
		V1Group: v1Group,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
