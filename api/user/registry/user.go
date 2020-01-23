package registry

import (
	"github.com/16francs/gran/api/user/internal/application"
	rv "github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain/service"
	"github.com/16francs/gran/api/user/internal/infrastructure/persistence"
	dv "github.com/16francs/gran/api/user/internal/infrastructure/validation"
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// V1UserInjection - v1 User関連の依存関係を記載
func V1UserInjection(fa *authentication.Auth, fs *firestore.Firestore) v1.APIV1UserHandler {
	up := persistence.NewUserPersistence(fa, fs)
	udv := dv.NewUserDomainValidation(up)
	us := service.NewUserService(udv, up)
	urv := rv.NewUserRequestValidation()
	uu := application.NewUserApplication(urv, us)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
