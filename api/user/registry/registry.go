package registry

import (
	"github.com/16francs/gran/api/user/internal/application/usecase"
	"github.com/16francs/gran/api/user/internal/infrastructure/persistence"
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	V1Health v1.APIV1HealthHandler
	V1User   v1.APIV1UserHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fb *firebase.Firebase, fs *firestore.Firestore) *Registry {
	v1Health := v1HealthInjection()
	v1User := v1UserInjection(*fs)

	return &Registry{
		V1User:   v1User,
		V1Health: v1Health,
	}
}

func v1HealthInjection() v1.APIV1HealthHandler {
	hh := v1.NewAPIV1HealthHandler()

	return hh
}

func v1UserInjection(fs firestore.Firestore) v1.APIV1UserHandler {
	up := persistence.NewUserPersistence(fs)
	uu := usecase.NewUserUsecase(up)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
