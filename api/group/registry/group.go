package registry

import (
	"github.com/16francs/gran/api/group/internal/application"
	rv "github.com/16francs/gran/api/group/internal/application/validation"
	"github.com/16francs/gran/api/group/internal/domain/service"
	"github.com/16francs/gran/api/group/internal/infrastructure/persistence"
	dv "github.com/16francs/gran/api/group/internal/infrastructure/validation"
	v1 "github.com/16francs/gran/api/group/internal/interface/handler/v1"
	"github.com/16francs/gran/api/group/lib/firebase/authentication"
	"github.com/16francs/gran/api/group/lib/firebase/firestore"
)

// V1GroupInjection - v1 Group関連の依存関係を記載
func V1GroupInjection(fa *authentication.Auth, fs *firestore.Firestore) v1.APIV1GroupHandler {
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
