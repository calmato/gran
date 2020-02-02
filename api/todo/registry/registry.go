package registry

import (
	"fmt"

	"github.com/16francs/gran/api/todo/lib/firebase/authentication"
	"github.com/16francs/gran/api/todo/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct{}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	fmt.Println(fa, fs)

	return &Registry{}
}
