# Application層

Interfaces層から情報を受け取り、Domain層で定義してある関数を用いて任意のビジネスロジックを実行する。

## テンプレート

```go
package application

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/calmato/gran/api/sample/internal/application/request"
	"github.com/calmato/gran/api/sample/internal/application/validation"
	"github.com/calmato/gran/api/sample/internal/domain"
	"github.com/calmato/gran/api/sample/internal/domain/service"
)

// SampleApplication - SampleApplicationインターフェース
type SampleApplication interface {
	Create(ctx context.Context, req *request.CreateSample) error
}

type sampleApplication struct {
	sampleRequestValidation validation.SampleRequestValidation
	sampleService           service.SampleService
}

// NewSampleApplication - SampleApplicationの生成
func NewSampleApplication(urv validation.SampleRequestValidation, us service.SampleService) SampleApplication {
	return &sampleApplication{
		sampleRequestValidation: urv,
		sampleService:           us,
	}
}

func (ua *sampleApplication) Create(ctx context.Context, req *request.CreateSample) error {
	if err := ua.sampleRequestValidation.CreateSample(req); err != nil {
		err = xerrors.Errorf("Failed to Application/RequestValidation: %w", err)
		return domain.InvalidRequestValidation.New(err)
	}

	current := time.Now()
	u := &domain.Sample{
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: current,
		UpdatedAt: current,
	}

	if err := ua.sampleService.Create(ctx, u); err != nil {
		return err
	}

	return nil
}
```
