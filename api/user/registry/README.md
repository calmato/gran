# Registry

Dependency Injection(DI)の実装。  
各層で実装している構造体の依存解決を解決する。

## テンプレート

```go
package registry

import (
	"github.com/16francs/gran/api/sample/internal/application"
	rv "github.com/16francs/gran/api/sample/internal/application/validation"
	"github.com/16francs/gran/api/sample/internal/domain/service"
	"github.com/16francs/gran/api/sample/internal/infrastructure/persistence"
	dv "github.com/16francs/gran/api/sample/internal/infrastructure/validation"
	v1 "github.com/16francs/gran/api/sample/internal/interface/handler/v1"
	"github.com/16francs/gran/api/sample/lib/firebase/authentication"
	"github.com/16francs/gran/api/sample/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	V1Health v1.APIV1HealthHandler
	V1Sample   v1.APIV1SampleHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	v1Health := v1HealthInjection()
	v1Sample := v1SampleInjection(fa, fs)

	return &Registry{
		V1Sample:   v1Sample,
		V1Health: v1Health,
	}
}

func v1HealthInjection() v1.APIV1HealthHandler {
	hh := v1.NewAPIV1HealthHandler()

	return hh
}

func v1SampleInjection(fa *authentication.Auth, fs *firestore.Firestore) v1.APIV1SampleHandler {
	sp := persistence.NewSamplePersistence(fa, fs)
	sdv := dv.NewSampleDomainValidation(sp)
	ss := service.NewSampleService(sdv, sp)
	srv := rv.NewSampleRequestValidation()
	su := application.NewSampleApplication(srv, ss)
	sh := v1.NewAPIV1SampleHandler(su)

	return uh
}
```
