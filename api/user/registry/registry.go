package registry

import (
	"github.com/16francs/gran/api/user/internal/interface/handler"
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	Health handler.APIHealthHandler
	V1User v1.APIV1UserHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	health := healthInjection()
	v1User := V1UserInjection(fa, fs)

	return &Registry{
		Health: health,
		V1User: v1User,
	}
}

func healthInjection() handler.APIHealthHandler {
	hh := handler.NewAPIHealthHandler()

	return hh
}
