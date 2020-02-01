package registry

import (
	"github.com/16francs/gran/api/group/internal/interface/handler"
	"github.com/16francs/gran/api/group/lib/firebase/authentication"
	"github.com/16francs/gran/api/group/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	Health handler.APIHealthHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	health := healthInjection()

	return &Registry{
		Health: health,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
