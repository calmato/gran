package registry

import (
	"fmt"

	"github.com/16francs/gran/api/todo/internal/interface/handler"
	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	Health handler.APIHealthHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	health := healthInjection()

	fmt.Println(fa, fs)

	return &Registry{
		Health: health,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
