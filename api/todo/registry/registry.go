package registry

import (
	"github.com/16francs/gran/api/todo/internal/interface/handler"
	v1 "github.com/16francs/gran/api/todo/internal/interface/handler/v1"
	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	Health  handler.APIHealthHandler
	V1Board v1.APIV1BoardHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	health := healthInjection()
	v1Board := V1BoardInjection(fa, fs)

	return &Registry{
		Health:  health,
		V1Board: v1Board,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
