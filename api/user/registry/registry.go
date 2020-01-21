package registry

import (
	"github.com/16francs/gran/api/user/internal/interface/handler"
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	V1Health handler.APIV1HealthHandler
	V1Group  v1.APIV1GroupHandler
	V1User   v1.APIV1UserHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	v1Health := v1HealthInjection()
	v1Group := V1GroupInjection(fa, fs)
	v1User := V1UserInjection(fa, fs)

	return &Registry{
		V1Group:  v1Group,
		V1Health: v1Health,
		V1User:   v1User,
	}
}

func v1HealthInjection() handler.APIV1HealthHandler {
	hh := handler.NewAPIV1HealthHandler()

	return hh
}
