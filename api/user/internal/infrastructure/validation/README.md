# Validation Infrastructure層

Domain層のValidationで定義したインターフェースを実装する。
Repositoryへのアクセスが必要な箇所は、別途バリデーションを実装する。

## テンプレート

```go
package validation

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/sample/internal/domain"
	"github.com/16francs/gran/api/sample/internal/domain/repository"
	dv "github.com/16francs/gran/api/sample/internal/domain/validation"
)

type sampleDomainValidation struct {
	validator      DomainValidator
	sampleRepository repository.SampleRepository
}

// NewSampleDomainValidation - SampleDomainValidationの生成
func NewSampleDomainValidation(ur repository.SampleRepository) dv.SampleDomainValidation {
	v := NewDomainValidator()

	return &sampleDomainValidation{
		validator:      v,
		sampleRepository: ur,
	}
}

func (udv *sampleDomainValidation) Sample(ctx context.Context, u *domain.Sample) error {
	err := udv.validator.Run(u)
	if err != nil {
		return err
	}

	err = uniqueCheckEmail(ctx, udv.sampleRepository, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func uniqueCheckEmail(ctx context.Context, ur repository.SampleRepository, email string) error {
	uid, _ := ur.GetUIDByEmail(ctx, email)
	if uid != "" {
		return xerrors.New("Email is not unique.")
	}

	return nil
}
```
