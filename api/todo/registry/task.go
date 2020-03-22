package registry

import (
	"github.com/16francs/gran/api/todo/internal/application"
	rv "github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain/service"
	"github.com/16francs/gran/api/todo/internal/infrastructure/persistence"
	dv "github.com/16francs/gran/api/todo/internal/infrastructure/validation"
	v1 "github.com/16francs/gran/api/todo/internal/interface/handler/v1"
	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
	gcs "github.com/16francs/gran/api/todo/lib/firebase/storage"
)

// V1TaskInjection - v1 Task関連の依存関係を記載
func V1TaskInjection(fa *authentication.Auth, fs *firestore.Firestore, _ *gcs.Storage) v1.APIV1TaskHandler {
	// fu := storage.NewFileUploader(cs)

	up := persistence.NewUserPersistence(fa, fs)
	us := service.NewUserService(up)

	tp := persistence.NewTaskPersistence(fs)
	tdv := dv.NewTaskDomainValidation()
	ts := service.NewTaskService(tdv, tp)
	trv := rv.NewTaskRequestValidation()
	ta := application.NewTaskApplication(trv, ts, us)
	th := v1.NewAPIV1TaskHandler(ta)

	return th
}
