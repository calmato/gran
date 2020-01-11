# Service Domain層

レポジトリへのアクセス、ドメインに関するバリデーションの処理を行う。

## テンプレート

```go
package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/sample/internal/domain/repository"

	"github.com/16francs/gran/api/sample/internal/domain"
	"github.com/16francs/gran/api/sample/internal/domain/validation"
)

// SampleService - SampleServiceインターフェース
type SampleService interface {
	Create(ctx context.Context, u *domain.Sample) error
}

type sampleService struct {
	sampleDomainValidation validation.SampleDomainValidation
	sampleRepository       repository.SampleRepository
}

// NewSampleService - SampleServiceの生成
func NewSampleService(udv validation.SampleDomainValidation, ur repository.SampleRepository) SampleService {
	return &sampleService{
		sampleDomainValidation: udv,
		sampleRepository:       ur,
	}
}

func (us *sampleService) Create(ctx context.Context, u *domain.Sample) error {
	if err := us.sampleDomainValidation.Sample(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/DomainValidation: %w", err)
		return domain.InvalidDomainValidation.New(err)
	}

	if err := us.sampleRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.Unknown.New(err)
	}

	return nil
}
```
