package registry

import (
	"github.com/16francs/gran/api/todo/internal/application"
	rv "github.com/16francs/gran/api/todo/internal/application/validation"
	"github.com/16francs/gran/api/todo/internal/domain/service"
	"github.com/16francs/gran/api/todo/internal/infrastructure/persistence"
	"github.com/16francs/gran/api/todo/internal/infrastructure/storage"
	dv "github.com/16francs/gran/api/todo/internal/infrastructure/validation"
	v1 "github.com/16francs/gran/api/todo/internal/interface/handler/v1"
	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
	gcs "github.com/16francs/gran/api/todo/lib/firebase/storage"
)

// V1BoardInjection - v1 Board関連の依存関係を記載
func V1BoardInjection(fa *authentication.Auth, fs *firestore.Firestore, cs *gcs.Storage) v1.APIV1BoardHandler {
	fu := storage.NewFileUploader(cs)

	up := persistence.NewUserPersistence(fa, fs)
	us := service.NewUserService(up)

	bp := persistence.NewBoardPersistence(fs)
	bdv := dv.NewBoardDomainValidation()
	bs := service.NewBoardService(bdv, bp, fu)
	brv := rv.NewBoardRequestValidation()
	bu := application.NewBoardApplication(brv, bs, us)
	bh := v1.NewAPIV1BoardHandler(bu)

	return bh
}
