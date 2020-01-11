# Domain Validation層

Domainのバリデーションを行う。  
バリデーション範囲は、文字数、正規表現等のエンティティ(データベース)に関連する内容。
処理の詳細は、Infrastructure層で実装するのでインターフェースのみ記述。

## テンプレート

```go
package validation

import (
	"context"

	"github.com/16francs/gran/api/sample/internal/domain"
)

// SampleDomainValidation - SampleDomainRepositoryインターフェース
type SampleDomainValidation interface {
	Sample(ctx context.Context, u *domain.Sample) error
}
```
