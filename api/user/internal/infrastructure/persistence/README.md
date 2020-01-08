# Persistence Infrastructure層

Domain層のRepositoryで定義したインタフェースを実装する。  
今回はDBとしてFirebase Authentication, Firestoreを使用する。

## テンプレート

```go
package persistence

import (
	"context"

	"github.com/16francs/gran/api/sample/internal/domain"
	"github.com/16francs/gran/api/sample/internal/domain/repository"
	"github.com/16francs/gran/api/sample/lib/firebase/authentication"
	"github.com/16francs/gran/api/sample/lib/firebase/firestore"
)

type samplePersistence struct {
	auth      *authentication.Auth
	firestore *firestore.Firestore
}

// SampleCollection - SampleCollection名
const SampleCollection = "samples"

// NewSamplePersistence - SampleRepositoryの生成
func NewSamplePersistence(fa *authentication.Auth, fs *firestore.Firestore) repository.SampleRepository {
	return &samplePersistence{
		auth:      fa,
		firestore: fs,
	}
}

func (r *samplePersistence) Create(ctx context.Context, u *domain.Sample) error {
	uid, err := createSampleInAuth(ctx, r.auth, u)
	if err != nil {
		return err
	}

	u.ID = uid

	if err = setInFirestore(ctx, r.firestore, u); err != nil {
		return err
	}

	return nil
}

func (r *samplePersistence) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := getUIDByEmailInAuth(ctx, r.auth, email)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func getUIDByEmailInAuth(ctx context.Context, fa *authentication.Auth, email string) (string, error) {
	return fa.GetUIDByEmail(ctx, email)
}

func createSampleInAuth(ctx context.Context, fa *authentication.Auth, u *domain.Sample) (string, error) {
	return fa.CreateSample(ctx, u.Email, u.Password)
}

func setInFirestore(ctx context.Context, fs *firestore.Firestore, u *domain.Sample) error {
	return fs.Set(ctx, SampleCollection, u.ID, u)
}
```
