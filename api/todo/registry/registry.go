package registry

import (
	"github.com/16francs/gran/api/todo/internal/interface/handler"
	v1 "github.com/16francs/gran/api/todo/internal/interface/handler/v1"
	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
	"github.com/16francs/gran/api/todo/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	Health  handler.APIHealthHandler
	V1Group v1.APIV1GroupHandler
	V1Board v1.APIV1BoardHandler
	V1Task  v1.APIV1TaskHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore, cs *storage.Storage) *Registry {
	health := healthInjection()
	v1Group := V1GroupInjection(fa, fs, cs)
	v1Board := V1BoardInjection(fa, fs, cs)
	v1Task := V1TaskInjection(fa, fs, cs)

	return &Registry{
		Health:  health,
		V1Group: v1Group,
		V1Board: v1Board,
		V1Task:  v1Task,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
