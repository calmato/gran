package persistence

import "strings"

const (
	// UserCollection - UserCollection名
	UserCollection = "users"
	// GroupCollection - GroupCollection名
	GroupCollection = "groups"
)

func getUserReference(userID string) string {
	return strings.Join([]string{UserCollection, userID}, "/")
}
