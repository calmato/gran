package domain

import "time"

// User - Userエンティティ
type User struct {
	ID           string    `firestore:"id"`
	Email        string    `firestore:"email" validate:"required,email,max=256"`
	Password     string    `firestore:"-" validate:"password,max=32"` // TODO: minも足したいので、カスタムバリデーション作成する
	Name         string    `firestore:"name"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	GroupRefs    []string  `firestore:"group_refs,omitempty"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
