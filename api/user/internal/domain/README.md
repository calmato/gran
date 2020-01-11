# Domain層

Domain層はドメインロジックを実装する責務を持つ。  
Domain層はDBアクセスなどの技術的な実装は持たず、それらはInfrastructure層が担当します。  
Domain層はどの層にも依存しないのでこの層だけで完結します。

## テンプレート

```go
package domain

import "time"

// Sample - Sampleエンティティ
type Sample struct {
	ID           string    `firestore:"id"`
	Email        string    `firestore:"email" validate:"email,max=256"`
	Password     string    `firestore:"-" validate:"password,min=6,max=32"`
	Name         string    `firestore:"name" validate:"max=32"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
```
