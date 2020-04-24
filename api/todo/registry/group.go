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
	"github.com/16francs/gran/api/todo/lib/firebase/storage"
)

// V1GroupInjection - v1 Group関連の依存関係を記載
func V1GroupInjection(fa *authentication.Auth, fs *firestore.Firestore, _ *storage.Storage) v1.APIV1GroupHandler {
	up := persistence.NewUserPersistence(fa, fs)
	us := service.NewUserService(up)

	bp := persistence.NewBoardPersistence(fs)

	gp := persistence.NewGroupPersistence(fs)
	gdv := dv.NewGroupDomainValidation()
	gs := service.NewGroupService(gdv, gp, up, bp)
	grv := rv.NewGroupRequestValidation()
	gu := application.NewGroupApplication(grv, gs, us)
	gh := v1.NewAPIV1GroupHandler(gu)

	return gh
}
