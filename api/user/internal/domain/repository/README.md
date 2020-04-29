# Repository Domain層

DBの処理を記述するインターフェースを作成する。  
処理の詳細は記述しない。

## テンプレート

```go
package repository

import (
	"context"

	"github.com/calmato/gran/api/sample/internal/domain"
)

// SampleRepository - SampleRepositoryインターフェース
type SampleRepository interface {
	Create(ctx context.Context, u *domain.Sample) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
```
