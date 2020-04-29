package registry

import (
	"github.com/calmato/gran/api/user/internal/application"
	rv "github.com/calmato/gran/api/user/internal/application/validation"
	"github.com/calmato/gran/api/user/internal/domain/service"
	"github.com/calmato/gran/api/user/internal/infrastructure/persistence"
	"github.com/calmato/gran/api/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran/api/user/internal/infrastructure/validation"
	v1 "github.com/calmato/gran/api/user/internal/interface/handler/v1"
	"github.com/calmato/gran/api/user/lib/firebase/authentication"
	"github.com/calmato/gran/api/user/lib/firebase/firestore"
	gcs "github.com/calmato/gran/api/user/lib/firebase/storage"
)

// V1UserInjection - v1 User関連の依存関係を記載
func V1UserInjection(fa *authentication.Auth, fs *firestore.Firestore, cs *gcs.Storage) v1.APIV1UserHandler {
	fu := storage.NewFileUploader(cs)

	up := persistence.NewUserPersistence(fa, fs)
	udv := dv.NewUserDomainValidation(up)
	us := service.NewUserService(udv, up, fu)
	urv := rv.NewUserRequestValidation()
	uu := application.NewUserApplication(urv, us)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
