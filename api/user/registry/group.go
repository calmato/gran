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

// V1GroupInjection - v1 Group関連の依存関係を記載
func V1GroupInjection(fa *authentication.Auth, fs *firestore.Firestore) v1.APIV1GroupHandler {
	up := persistence.NewUserPersistence(fa, fs)
	udv := dv.NewUserDomainValidation(up)
	us := service.NewUserService(udv, up)

	gp := persistence.NewGroupPersistence(fs)
	gdv := dv.NewGroupDomainValidation(gp)
	gs := service.NewGroupService(gdv, gp)
	grv := rv.NewGroupRequestValidation()
	gu := application.NewGroupApplication(grv, gs, us)
	gh := v1.NewAPIV1GroupHandler(gu)

	return gh
}
