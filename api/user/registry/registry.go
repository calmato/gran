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
	firebase.Firebase
	firestore.Firestore
	v1.APIV1HealthHandler
	v1.APIV1UserHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fb *firebase.Firebase, fs *firestore.Firestore) *Registry {
	apiV1HealthHandler := v1HealthInjection()
	apiV1UserHandler := v1UserInjection(*fs)

	return &Registry{
		Firebase:           *fb,
		Firestore:          *fs,
		APIV1HealthHandler: apiV1HealthHandler,
		APIV1UserHandler:   apiV1UserHandler,
	}
}

func v1HealthInjection() v1.APIV1HealthHandler {
	apiV1HealthHandler := v1.NewAPIV1HealthHandler()

	return apiV1HealthHandler
}

func v1UserInjection(fs firestore.Firestore) v1.APIV1UserHandler {
	up := persistence.NewUserPersistence(fs)
	uu := usecase.NewUserUsecase(up)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
