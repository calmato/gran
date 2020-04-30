# Request Validation層

Request値のバリデーションを行う。  
バリデーション範囲は、存在性、同一性等のフォームに関連する内容。

## テンプレート

```go
package validation

import "github.com/calmato/gran/api/sample/internal/application/request"

// SampleRequestValidation - ユーザー関連のバリデーション
type SampleRequestValidation interface {
	CreateSample(cu *request.CreateSample) error
}

type sampleRequestValidation struct {
	validator RequestValidator
}

// NewSampleRequestValidation - SampleRequestValidationの生成
func NewSampleRequestValidation() SampleRequestValidation {
	rv := NewRequestValidator()

	return &sampleRequestValidation{
		validator: rv,
	}
}

func (urv *sampleRequestValidation) CreateSample(cu *request.CreateSample) error {
	return urv.validator.Run(cu)
}
```
