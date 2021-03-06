package registry

import (
	"github.com/calmato/gran/api/todo/internal/application"
	rv "github.com/calmato/gran/api/todo/internal/application/validation"
	"github.com/calmato/gran/api/todo/internal/domain/service"
	"github.com/calmato/gran/api/todo/internal/infrastructure/persistence"
	"github.com/calmato/gran/api/todo/internal/infrastructure/storage"
	dv "github.com/calmato/gran/api/todo/internal/infrastructure/validation"
	v1 "github.com/calmato/gran/api/todo/internal/interface/handler/v1"
	"github.com/calmato/gran/api/todo/lib/firebase/authentication"
	"github.com/calmato/gran/api/todo/lib/firebase/firestore"
	gcs "github.com/calmato/gran/api/todo/lib/firebase/storage"
)

// V1TaskInjection - v1 Task関連の依存関係を記載
func V1TaskInjection(fa *authentication.Auth, fs *firestore.Firestore, cs *gcs.Storage) v1.APIV1TaskHandler {
	fu := storage.NewFileUploader(cs)

	up := persistence.NewUserPersistence(fa, fs)
	us := service.NewUserService(up)

	bp := persistence.NewBoardPersistence(fs)
	bdv := dv.NewBoardDomainValidation()

	tp := persistence.NewTaskPersistence(fs)
	tdv := dv.NewTaskDomainValidation()

	bs := service.NewBoardService(bdv, bp, tp, fu)

	ts := service.NewTaskService(tdv, tp, bp)
	trv := rv.NewTaskRequestValidation()
	ta := application.NewTaskApplication(trv, ts, bs, us)
	th := v1.NewAPIV1TaskHandler(ta)

	return th
}
